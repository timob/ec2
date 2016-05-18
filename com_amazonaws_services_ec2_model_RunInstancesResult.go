package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRunInstancesResultInterface interface {
	JavaLangObjectInterface

	// public void setReservation(com.amazonaws.services.ec2.model.Reservation)
	SetReservation(a ServicesEc2ModelReservationInterface) 

	// public com.amazonaws.services.ec2.model.Reservation getReservation()
	GetReservation() *ServicesEc2ModelReservation

	// public com.amazonaws.services.ec2.model.RunInstancesResult withReservation(com.amazonaws.services.ec2.model.Reservation)
	WithReservation(a ServicesEc2ModelReservationInterface) *ServicesEc2ModelRunInstancesResult

	// public com.amazonaws.services.ec2.model.RunInstancesResult clone()
	Clone() *ServicesEc2ModelRunInstancesResult
}

type ServicesEc2ModelRunInstancesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.RunInstancesResult()
func NewServicesEc2ModelRunInstancesResult() (*ServicesEc2ModelRunInstancesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RunInstancesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRunInstancesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setReservation(com.amazonaws.services.ec2.model.Reservation)
func (jbobject *ServicesEc2ModelRunInstancesResult) SetReservation(a ServicesEc2ModelReservationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservation", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Reservation"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Reservation getReservation()
func (jbobject *ServicesEc2ModelRunInstancesResult) GetReservation() *ServicesEc2ModelReservation {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservation", "com/amazonaws/services/ec2/model/Reservation")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelReservation{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RunInstancesResult withReservation(com.amazonaws.services.ec2.model.Reservation)
func (jbobject *ServicesEc2ModelRunInstancesResult) WithReservation(a ServicesEc2ModelReservationInterface) *ServicesEc2ModelRunInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservation", "com/amazonaws/services/ec2/model/RunInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Reservation"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelRunInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRunInstancesResult) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean equals(java.lang.Object)
func (jbobject *ServicesEc2ModelRunInstancesResult) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *ServicesEc2ModelRunInstancesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.RunInstancesResult clone()
func (jbobject *ServicesEc2ModelRunInstancesResult) Clone() *ServicesEc2ModelRunInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/RunInstancesResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelRunInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelRunInstancesResult) Clone2() (*JavaLangObject, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
	if err != nil {
		var zero *JavaLangObject
		return zero, err
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x, nil
}


