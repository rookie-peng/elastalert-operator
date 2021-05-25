package e2e

import (
	"context"
	"elastalert/api/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"time"
)

const WaitForStabilityTimeout = time.Minute * 4
const interval = time.Second * 1
const timeout = time.Second * 30

var _ = Describe("Elastalert Controller", func() {
	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
	})

	AfterEach(func() {
		key := types.NamespacedName{
			Name:      "e2e-elastalert",
			Namespace: "default",
		}
		ea := &v1alpha1.Elastalert{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: key.Namespace,
				Name:      key.Name,
			},
		}
		_ = k8sClient.Delete(context.Background(), ea)
		// Add any teardown steps that needs to be executed after each test
	})
	Context("Deploy Elastalert", func() {
		It("Test create Elastalert", func() {
			key := types.NamespacedName{
				Name:      "e2e-elastalert",
				Namespace: "default",
			}
			elastalert := &v1alpha1.Elastalert{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: key.Namespace,
					Name:      key.Name,
				},
				Spec: v1alpha1.ElastalertSpec{
					ConfigSetting: v1alpha1.NewFreeForm(map[string]interface{}{
						"es_host":         "es.com.cn",
						"es_port":         9200,
						"use_ssl":         true,
						"es_username":     "elastic",
						"es_password":     "changeme",
						"verify_certs":    false,
						"writeback_index": "elastalert",
						"run_every": map[string]interface{}{
							"minutes": 1,
						},
						"buffer_time": map[string]interface{}{
							"minutes": 15,
						},
					}),
					Rule: []v1alpha1.FreeForm{
						v1alpha1.NewFreeForm(map[string]interface{}{
							"name":  "test-elastalert",
							"type":  "any",
							"index": "api-*",
							"filter": []map[string]interface{}{
								{
									"query": map[string]interface{}{
										"query_string": map[string]interface{}{
											"query": "http_status_code: 503",
										},
									},
								},
							},
						}),
					},
					Alert: v1alpha1.NewFreeForm(map[string]interface{}{

						"alert": []string{
							"post",
						},
						"http_post_url":     "https://api-b2b-prd.cn-pgcloud.com/nymeria/alerts",
						"http_post_timeout": 60,
					}),
					PodTemplateSpec: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Annotations: map[string]string{
								"e2e": "test",
							},
						},
						Spec: v1.PodSpec{
							Containers: []v1.Container{},
						},
					},
				},
			}
			Expect(k8sClient.Create(context.Background(), elastalert)).Should(Succeed())
			Eventually(func() error {
				dep := &appsv1.Deployment{}
				err := k8sClient.Get(context.Background(), key, dep)
				return err
			}, timeout, interval).Should(Succeed())
		})
		It("Test create Elastalert with wrong config", func() {
			key := types.NamespacedName{
				Name:      "e2e-elastalert",
				Namespace: "default",
			}
			elastalert := &v1alpha1.Elastalert{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: key.Namespace,
					Name:      key.Name,
				},
				Spec: v1alpha1.ElastalertSpec{
					ConfigSetting: v1alpha1.NewFreeForm(map[string]interface{}{
						"config": "test",
					}),
					Rule: []v1alpha1.FreeForm{
						v1alpha1.NewFreeForm(map[string]interface{}{
							"name": "test-elastalert",
							"type": "any",
						}),
					},
					PodTemplateSpec: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{},
						},
					},
				},
			}
			Expect(k8sClient.Create(context.Background(), elastalert)).Should(Succeed())
			By("Start waiting for failed status")
			Eventually(func() string {
				ea := &v1alpha1.Elastalert{}
				_ = k8sClient.Get(context.Background(), key, ea)
				return ea.Status.Phase
			}, WaitForStabilityTimeout, interval).Should(Equal("FAILED"))
		})
	})
})
