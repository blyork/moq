// Code generated by moq; DO NOT EDIT.
// github.com/blyork/moq

package generate

import (
	"sync"
)

// Ensure, that MyInterfaceMock does implement MyInterface.
// If this is not the case, regenerate this file with moq.
var _ MyInterface = &MyInterfaceMock{}

// MyInterfaceMock is a mock implementation of MyInterface.
//
//     func TestSomethingThatUsesMyInterface(t *testing.T) {
//
//         // make and configure a mocked MyInterface
//         mockedMyInterface := &MyInterfaceMock{
//             OneFunc: func() bool {
// 	               panic("moqInstance out the One method")
//             },
//             ThreeFunc: func() string {
// 	               panic("moqInstance out the Three method")
//             },
//             TwoFunc: func() int {
// 	               panic("moqInstance out the Two method")
//             },
//         }
//
//         // use mockedMyInterface in code that requires MyInterface
//         // and then make assertions.
//
//     }
type MyInterfaceMock struct {
	// OneFunc mocks the One method.
	OneFunc func() bool

	// ThreeFunc mocks the Three method.
	ThreeFunc func() string

	// TwoFunc mocks the Two method.
	TwoFunc func() int

	// calls tracks calls to the methods.
	calls struct {
		// One holds details about calls to the One method.
		One struct {
			Calls []MyInterfaceMockOneCall
			Mu    sync.RWMutex
		}
		// Three holds details about calls to the Three method.
		Three struct {
			Calls []MyInterfaceMockThreeCall
			Mu    sync.RWMutex
		}
		// Two holds details about calls to the Two method.
		Two struct {
			Calls []MyInterfaceMockTwoCall
			Mu    sync.RWMutex
		}
	}
}

// One calls OneFunc.
func (moqInstance *MyInterfaceMock) One() bool {
	if moqInstance.OneFunc == nil {
		panic("MyInterfaceMock.OneFunc: method is nil but MyInterface.One was just called")
	}
	var moqCallIndex int
	moqInstance.calls.One.Mu.Lock()
	moqCallIndex = len(moqInstance.calls.One.Calls)
	moqInstance.calls.One.Calls = append(moqInstance.calls.One.Calls, MyInterfaceMockOneCall{
		In: MyInterfaceMockOneCallInput{},
	})
	moqInstance.calls.One.Mu.Unlock()
	out1 := moqInstance.OneFunc()
	moqInstance.calls.One.Mu.Lock()
	moqInstance.calls.One.Calls[moqCallIndex].Out = MyInterfaceMockOneCallOutput{
		Out1: out1,
	}
	moqInstance.calls.One.Mu.Unlock()
}

// MyInterfaceMockOneCallInput represents the input to a MyInterfaceMock.One call.
type MyInterfaceMockOneCallInput struct {
}

// MyInterfaceMockOneCallOutput represents the output for a MyInterfaceMock.One call.
type MyInterfaceMockOneCallOutput struct {
	Out1 bool
}

// MyInterfaceMockOneCall represents the input and output for a MyInterfaceMock.One call.
type MyInterfaceMockOneCall struct {
	In  MyInterfaceMockOneCallInput
	Out MyInterfaceMockOneCallOutput
}

// OneCalls gets all the calls that were made to One.
// Check the length with:
//     len(mockedMyInterface.OneCalls())
func (moqInstance *MyInterfaceMock) OneCalls() []MyInterfaceMockOneCall {
	var calls []MyInterfaceMockOneCall
	moqInstance.calls.One.Mu.RLock()
	calls = moqInstance.calls.One.Calls
	moqInstance.calls.One.Mu.RUnlock()
	return calls
}

// Three calls ThreeFunc.
func (moqInstance *MyInterfaceMock) Three() string {
	if moqInstance.ThreeFunc == nil {
		panic("MyInterfaceMock.ThreeFunc: method is nil but MyInterface.Three was just called")
	}
	var moqCallIndex int
	moqInstance.calls.Three.Mu.Lock()
	moqCallIndex = len(moqInstance.calls.Three.Calls)
	moqInstance.calls.Three.Calls = append(moqInstance.calls.Three.Calls, MyInterfaceMockThreeCall{
		In: MyInterfaceMockThreeCallInput{},
	})
	moqInstance.calls.Three.Mu.Unlock()
	out1 := moqInstance.ThreeFunc()
	moqInstance.calls.Three.Mu.Lock()
	moqInstance.calls.Three.Calls[moqCallIndex].Out = MyInterfaceMockThreeCallOutput{
		Out1: out1,
	}
	moqInstance.calls.Three.Mu.Unlock()
}

// MyInterfaceMockThreeCallInput represents the input to a MyInterfaceMock.Three call.
type MyInterfaceMockThreeCallInput struct {
}

// MyInterfaceMockThreeCallOutput represents the output for a MyInterfaceMock.Three call.
type MyInterfaceMockThreeCallOutput struct {
	Out1 string
}

// MyInterfaceMockThreeCall represents the input and output for a MyInterfaceMock.Three call.
type MyInterfaceMockThreeCall struct {
	In  MyInterfaceMockThreeCallInput
	Out MyInterfaceMockThreeCallOutput
}

// ThreeCalls gets all the calls that were made to Three.
// Check the length with:
//     len(mockedMyInterface.ThreeCalls())
func (moqInstance *MyInterfaceMock) ThreeCalls() []MyInterfaceMockThreeCall {
	var calls []MyInterfaceMockThreeCall
	moqInstance.calls.Three.Mu.RLock()
	calls = moqInstance.calls.Three.Calls
	moqInstance.calls.Three.Mu.RUnlock()
	return calls
}

// Two calls TwoFunc.
func (moqInstance *MyInterfaceMock) Two() int {
	if moqInstance.TwoFunc == nil {
		panic("MyInterfaceMock.TwoFunc: method is nil but MyInterface.Two was just called")
	}
	var moqCallIndex int
	moqInstance.calls.Two.Mu.Lock()
	moqCallIndex = len(moqInstance.calls.Two.Calls)
	moqInstance.calls.Two.Calls = append(moqInstance.calls.Two.Calls, MyInterfaceMockTwoCall{
		In: MyInterfaceMockTwoCallInput{},
	})
	moqInstance.calls.Two.Mu.Unlock()
	out1 := moqInstance.TwoFunc()
	moqInstance.calls.Two.Mu.Lock()
	moqInstance.calls.Two.Calls[moqCallIndex].Out = MyInterfaceMockTwoCallOutput{
		Out1: out1,
	}
	moqInstance.calls.Two.Mu.Unlock()
}

// MyInterfaceMockTwoCallInput represents the input to a MyInterfaceMock.Two call.
type MyInterfaceMockTwoCallInput struct {
}

// MyInterfaceMockTwoCallOutput represents the output for a MyInterfaceMock.Two call.
type MyInterfaceMockTwoCallOutput struct {
	Out1 int
}

// MyInterfaceMockTwoCall represents the input and output for a MyInterfaceMock.Two call.
type MyInterfaceMockTwoCall struct {
	In  MyInterfaceMockTwoCallInput
	Out MyInterfaceMockTwoCallOutput
}

// TwoCalls gets all the calls that were made to Two.
// Check the length with:
//     len(mockedMyInterface.TwoCalls())
func (moqInstance *MyInterfaceMock) TwoCalls() []MyInterfaceMockTwoCall {
	var calls []MyInterfaceMockTwoCall
	moqInstance.calls.Two.Mu.RLock()
	calls = moqInstance.calls.Two.Calls
	moqInstance.calls.Two.Mu.RUnlock()
	return calls
}
