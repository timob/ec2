package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateSpotDatafeedSubscriptionResultInterface interface {
	JavaLangObjectInterface

	// public void setSpotDatafeedSubscription(com.amazonaws.services.ec2.model.SpotDatafeedSubscription)
	SetSpotDatafeedSubscription(a ServicesEc2ModelSpotDatafeedSubscriptionInterface) 

	// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription getSpotDatafeedSubscription()
	GetSpotDatafeedSubscription() *ServicesEc2ModelSpotDatafeedSubscription

	// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionResult withSpotDatafeedSubscription(com.amazonaws.services.ec2.model.SpotDatafeedSubscription)
	WithSpotDatafeedSubscription(a ServicesEc2ModelSpotDatafeedSubscriptionInterface) *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult

	// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionResult clone()
	Clone() *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult
}

type ServicesEc2ModelCreateSpotDatafeedSubscriptionResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionResult()
func NewServicesEc2ModelCreateSpotDatafeedSubscriptionResult() (*ServicesEc2ModelCreateSpotDatafeedSubscriptionResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateSpotDatafeedSubscriptionResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateSpotDatafeedSubscriptionResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSpotDatafeedSubscription(com.amazonaws.services.ec2.model.SpotDatafeedSubscription)
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult) SetSpotDatafeedSubscription(a ServicesEc2ModelSpotDatafeedSubscriptionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotDatafeedSubscription", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotDatafeedSubscription"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription getSpotDatafeedSubscription()
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult) GetSpotDatafeedSubscription() *ServicesEc2ModelSpotDatafeedSubscription {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotDatafeedSubscription", "com/amazonaws/services/ec2/model/SpotDatafeedSubscription")
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
	unique_x := &ServicesEc2ModelSpotDatafeedSubscription{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionResult withSpotDatafeedSubscription(com.amazonaws.services.ec2.model.SpotDatafeedSubscription)
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult) WithSpotDatafeedSubscription(a ServicesEc2ModelSpotDatafeedSubscriptionInterface) *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotDatafeedSubscription", "com/amazonaws/services/ec2/model/CreateSpotDatafeedSubscriptionResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotDatafeedSubscription"))
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
	unique_x := &ServicesEc2ModelCreateSpotDatafeedSubscriptionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionResult clone()
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult) Clone() *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateSpotDatafeedSubscriptionResult")
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
	unique_x := &ServicesEc2ModelCreateSpotDatafeedSubscriptionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult) Clone2() (*JavaLangObject, error) {
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


