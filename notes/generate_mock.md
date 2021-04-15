# How to create Mocks

### Example

Create an interface:

    package doer
	type Doer interface {
	    DoSomething(int, string) error
	}

Create the implementation of the interface
 

    package user

	import "github.com/sgreben/testing-with-gomock/doer"

	type User struct {
	    Doer doer.Doer
	}

	func (u *User) Use() error {
	    return u.Doer.DoSomething(123, "Hello GoMock")
	}

Run the command to generate the mock for that interface:

    $GOPATH/bin/mockgen -destination=mocks/mock_doer.go -package=mocks github.com/sgreben/testing-with-gomock/doer Doer

where:

>  `-destination=mocks/mock_doer.go`: put the generated mocks in the file `mocks/mock_doer.go`.
   `-package=mocks`: put the generated mocks in the package `mocks`
   `github.com/sgreben/testing-with-gomock/doer`: generate mocks for this package
   `Doer`: generate mocks for this interface. This argument is required — we need to specify the interfaces to generate mocks for explicitly. We _can_, however specify multiple interfaces here as a comma-separated list (e.g. `Doer1,Doer2`).


To use the mock:

    mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)

To mock the call and the return:

	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)