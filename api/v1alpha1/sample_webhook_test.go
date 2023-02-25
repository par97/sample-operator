/* start_Copyright_Notice
Licensed Materials - Property of IBM

IBM Spectrum Fusion 5900-AOY
(C) Copyright IBM Corp. 2023 All Rights Reserved.

US Government Users Restricted Rights - Use, duplication or
disclosure restricted by GSA ADP Schedule Contract with
IBM Corp.
end_Copyright_Notice */

package v1alpha1

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	timeout  = time.Second * 30
	interval = time.Millisecond * 250
)

var _ = Describe("Test sample Webhook", func() {
	Context("Test create sample CR", func() {
		ctx := context.Background()
		// var err error

		BeforeEach(func() {
		})

		AfterEach(func() {
		})
		It("When create sample CR", func() {
			sample := &Sample{
				ObjectMeta: metav1.ObjectMeta{Name: "name1", Namespace: "default"},
				Spec:       SampleSpec{},
			}
			Expect(k8sClient.Create(ctx, sample)).Should(Succeed())
		})

		It("When create sample CR with forbidden_foo", func() {
			sample := &Sample{
				ObjectMeta: metav1.ObjectMeta{Name: "name2", Namespace: "default"},
				Spec: SampleSpec{
					Foo: "forbidden_foo",
				},
			}
			err := k8sClient.Create(ctx, sample)
			// fmt.Printf("err: %v\n", err)
			Expect(err).ShouldNot(Succeed())
			Expect(err.Error()).Should(ContainSubstring("not allowed"))
		})

	})
})
