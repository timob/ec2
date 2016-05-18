package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSpotDatafeedSubscriptionResultInterface interface {
	JavaLangObjectInterface

	// public void setSpotDatafeedSubscription(com.amazonaws.services.ec2.model.SpotDatafeedSubscription)
	SetSpotDatafeedSubscription(a ServicesEc2ModelSpotDatafeedSubscriptionInterface) 

	// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription getSpotDatafeedSubscription()
	GetSpotDatafeedSubscription() *ServicesEc2ModelSpotDatafeedSubscription

	// public com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionResult withSpotDatafeedSubscription(com.amazonaws.services.ec2.model.SpotDatafeedSubscription)
	WithSpotDatafeedSubscription(a ServicesEc2ModelSpotDatafeedSubscriptionInterface) *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult

	// public com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionResult clone()
	Clone() *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult
}

type ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionResult()
func NewServicesEc2ModelDescribeSpotDatafeedSubscriptionResult() (*ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSpotDatafeedSubscriptionResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSpotDatafeedSubscription(com.amazonaws.services.ec2.model.SpotDatafeedSubscription)
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult) SetSpotDatafeedSubscription(a ServicesEc2ModelSpotDatafeedSubscriptionInterface)  {
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
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult) GetSpotDatafeedSubscription() *ServicesEc2ModelSpotDatafeedSubscription {
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

// public com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionResult withSpotDatafeedSubscription(com.amazonaws.services.ec2.model.SpotDatafeedSubscription)
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult) WithSpotDatafeedSubscription(a ServicesEc2ModelSpotDatafeedSubscriptionInterface) *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotDatafeedSubscription", "com/amazonaws/services/ec2/model/DescribeSpotDatafeedSubscriptionResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotDatafeedSubscription"))
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
	unique_x := &ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionResult clone()
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult) Clone() *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSpotDatafeedSubscriptionResult")
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
	unique_x := &ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult) Clone2() (*JavaLangObject, error) {
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


