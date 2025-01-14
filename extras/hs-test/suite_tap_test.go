package main

import (
	"reflect"
	"runtime"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
)

type TapSuite struct {
	HstSuite
}

var tapTests = []func(s *TapSuite){}
var tapSoloTests = []func(s *TapSuite){}

func registerTapTests(tests ...func(s *TapSuite)) {
	tapTests = append(tapTests, tests...)
}
func registerTapSoloTests(tests ...func(s *TapSuite)) {
	tapSoloTests = append(tapSoloTests, tests...)
}

func (s *TapSuite) SetupSuite() {
	time.Sleep(1 * time.Second)
	s.HstSuite.SetupSuite()
	s.configureNetworkTopology("tap")
}

var _ = Describe("TapSuite", Ordered, ContinueOnFailure, func() {
	var s TapSuite
	BeforeAll(func() {
		s.SetupSuite()
	})
	BeforeEach(func() {
		s.SetupTest()
	})
	AfterAll(func() {
		s.TearDownSuite()
	})
	AfterEach(func() {
		s.TearDownTest()
	})

	for _, test := range tapTests {
		test := test
		pc := reflect.ValueOf(test).Pointer()
		funcValue := runtime.FuncForPC(pc)
		testName := strings.Split(funcValue.Name(), ".")[2]
		It(testName, func(ctx SpecContext) {
			s.log(testName + ": BEGIN")
			test(&s)
		}, SpecTimeout(suiteTimeout))
	}
})

var _ = Describe("TapSuiteSolo", Ordered, ContinueOnFailure, Serial, func() {
	var s TapSuite
	BeforeAll(func() {
		s.SetupSuite()
	})
	BeforeEach(func() {
		s.SetupTest()
	})
	AfterAll(func() {
		s.TearDownSuite()
	})
	AfterEach(func() {
		s.TearDownTest()
	})

	for _, test := range tapSoloTests {
		test := test
		pc := reflect.ValueOf(test).Pointer()
		funcValue := runtime.FuncForPC(pc)
		testName := strings.Split(funcValue.Name(), ".")[2]
		It(testName, Label("SOLO"), func(ctx SpecContext) {
			s.log(testName + ": BEGIN")
			test(&s)
		}, SpecTimeout(suiteTimeout))
	}
})
