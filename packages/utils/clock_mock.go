
// MockClock is an autogenerated mock type for the Clock type
type MockClock struct {
	mock.Mock
}

// Now provides a mock function with given fields:
func (_m *MockClock) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}