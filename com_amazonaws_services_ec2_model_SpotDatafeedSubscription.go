package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelSpotDatafeedSubscriptionInterface interface {
	JavaLangObjectInterface

	// public void setOwnerId(java.lang.String)
	SetOwnerId(a string) 

	// public java.lang.String getOwnerId()
	GetOwnerId() string

	// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription withOwnerId(java.lang.String)
	WithOwnerId(a string) *ServicesEc2ModelSpotDatafeedSubscription

	// public void setBucket(java.lang.String)
	SetBucket(a string) 

	// public java.lang.String getBucket()
	GetBucket() string

	// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription withBucket(java.lang.String)
	WithBucket(a string) *ServicesEc2ModelSpotDatafeedSubscription

	// public void setPrefix(java.lang.String)
	SetPrefix(a string) 

	// public java.lang.String getPrefix()
	GetPrefix() string

	// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription withPrefix(java.lang.String)
	WithPrefix(a string) *ServicesEc2ModelSpotDatafeedSubscription

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelSpotDatafeedSubscription

	// public void setState(com.amazonaws.services.ec2.model.DatafeedSubscriptionState)
	SetState(a ServicesEc2ModelDatafeedSubscriptionStateInterface) 

	// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription withState(com.amazonaws.services.ec2.model.DatafeedSubscriptionState)
	WithState(a ServicesEc2ModelDatafeedSubscriptionStateInterface) *ServicesEc2ModelSpotDatafeedSubscription

	// public void setFault(com.amazonaws.services.ec2.model.SpotInstanceStateFault)
	SetFault(a ServicesEc2ModelSpotInstanceStateFaultInterface) 

	// public com.amazonaws.services.ec2.model.SpotInstanceStateFault getFault()
	GetFault() *ServicesEc2ModelSpotInstanceStateFault

	// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription withFault(com.amazonaws.services.ec2.model.SpotInstanceStateFault)
	WithFault(a ServicesEc2ModelSpotInstanceStateFaultInterface) *ServicesEc2ModelSpotDatafeedSubscription

	// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription clone()
	Clone() *ServicesEc2ModelSpotDatafeedSubscription
}

type ServicesEc2ModelSpotDatafeedSubscription struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription()
func NewServicesEc2ModelSpotDatafeedSubscription() (*ServicesEc2ModelSpotDatafeedSubscription) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/SpotDatafeedSubscription")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelSpotDatafeedSubscription{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) SetOwnerId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOwnerId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getOwnerId()
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) GetOwnerId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOwnerId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription withOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) WithOwnerId(a string) *ServicesEc2ModelSpotDatafeedSubscription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOwnerId", "com/amazonaws/services/ec2/model/SpotDatafeedSubscription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotDatafeedSubscription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setBucket(java.lang.String)
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) SetBucket(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBucket", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getBucket()
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) GetBucket() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBucket", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription withBucket(java.lang.String)
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) WithBucket(a string) *ServicesEc2ModelSpotDatafeedSubscription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBucket", "com/amazonaws/services/ec2/model/SpotDatafeedSubscription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotDatafeedSubscription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrefix(java.lang.String)
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) SetPrefix(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrefix", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPrefix()
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) GetPrefix() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrefix", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription withPrefix(java.lang.String)
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) WithPrefix(a string) *ServicesEc2ModelSpotDatafeedSubscription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrefix", "com/amazonaws/services/ec2/model/SpotDatafeedSubscription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotDatafeedSubscription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) SetState2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getState()
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) GetState() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getState", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription withState(java.lang.String)
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) WithState2(a string) *ServicesEc2ModelSpotDatafeedSubscription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/SpotDatafeedSubscription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotDatafeedSubscription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.DatafeedSubscriptionState)
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) SetState(a ServicesEc2ModelDatafeedSubscriptionStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DatafeedSubscriptionState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription withState(com.amazonaws.services.ec2.model.DatafeedSubscriptionState)
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) WithState(a ServicesEc2ModelDatafeedSubscriptionStateInterface) *ServicesEc2ModelSpotDatafeedSubscription {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/SpotDatafeedSubscription", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DatafeedSubscriptionState"))
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
	unique_x := &ServicesEc2ModelSpotDatafeedSubscription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFault(com.amazonaws.services.ec2.model.SpotInstanceStateFault)
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) SetFault(a ServicesEc2ModelSpotInstanceStateFaultInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFault", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotInstanceStateFault"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotInstanceStateFault getFault()
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) GetFault() *ServicesEc2ModelSpotInstanceStateFault {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFault", "com/amazonaws/services/ec2/model/SpotInstanceStateFault")
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
	unique_x := &ServicesEc2ModelSpotInstanceStateFault{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription withFault(com.amazonaws.services.ec2.model.SpotInstanceStateFault)
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) WithFault(a ServicesEc2ModelSpotInstanceStateFaultInterface) *ServicesEc2ModelSpotDatafeedSubscription {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFault", "com/amazonaws/services/ec2/model/SpotDatafeedSubscription", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotInstanceStateFault"))
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
	unique_x := &ServicesEc2ModelSpotDatafeedSubscription{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) ToString() string {
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
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.SpotDatafeedSubscription clone()
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) Clone() *ServicesEc2ModelSpotDatafeedSubscription {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/SpotDatafeedSubscription")
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

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelSpotDatafeedSubscription) Clone2() (*JavaLangObject, error) {
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


