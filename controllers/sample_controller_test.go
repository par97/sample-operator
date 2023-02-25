package controllers

import (
	"context"
	"time"

	mygroupv1alpha1 "github.com/example/sample-operator/api/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("sample controller:", func() {

	// Define utility constants for object names and testing timeouts/durations and intervals.
	const (
		sampleName      = "sample1"
		sampleNamespace = "default"
		FooValue        = "foo1"

		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Second
	)

	Context("Test sample controller", func() {

		It("Should update sample status", func() {
			By("Create a new sample")
			ctx := context.Background()
			sample := &mygroupv1alpha1.Sample{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sampleName,
					Namespace: sampleNamespace,
				},
				Spec: mygroupv1alpha1.SampleSpec{
					Foo: FooValue,
				},
			}
			Expect(k8sClient.Create(ctx, sample)).Should(Succeed())

			By("Check the new sample status")
			Eventually(func() bool {
				err := k8sClient.Get(ctx, types.NamespacedName{Name: sampleName, Namespace: sampleNamespace}, sample)
				if err != nil {
					return false
				}
				return err == nil && sample.Status.State == FooValue
			}, timeout, interval).Should(BeTrue())
		})
	})
})
