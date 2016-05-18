package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceStatusSummaryInterface interface {
	JavaLangObjectInterface

	// public void setStatus(java.lang.String)
	SetStatus2(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.InstanceStatusSummary withStatus(java.lang.String)
	WithStatus2(a string) *ServicesEc2ModelInstanceStatusSummary

	// public void setStatus(com.amazonaws.services.ec2.model.SummaryStatus)
	SetStatus(a ServicesEc2ModelSummaryStatusInterface) 

	// public com.amazonaws.services.ec2.model.InstanceStatusSummary withStatus(com.amazonaws.services.ec2.model.SummaryStatus)
	WithStatus(a ServicesEc2ModelSummaryStatusInterface) *ServicesEc2ModelInstanceStatusSummary

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceStatusDetails> getDetails()
	GetDetails() []*ServicesEc2ModelInstanceStatusDetails

	// public void setDetails(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStatusDetails>)
	SetDetails(a []*ServicesEc2ModelInstanceStatusDetails) 

	// public com.amazonaws.services.ec2.model.InstanceStatusSummary withDetails(com.amazonaws.services.ec2.model.InstanceStatusDetails...)
	WithDetails(a ...*ServicesEc2ModelInstanceStatusDetails) *ServicesEc2ModelInstanceStatusSummary

	// public com.amazonaws.services.ec2.model.InstanceStatusSummary withDetails(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStatusDetails>)
	WithDetails2(a []*ServicesEc2ModelInstanceStatusDetails) *ServicesEc2ModelInstanceStatusSummary

	// public com.amazonaws.services.ec2.model.InstanceStatusSummary clone()
	Clone() *ServicesEc2ModelInstanceStatusSummary
}

type ServicesEc2ModelInstanceStatusSummary struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceStatusSummary()
func NewServicesEc2ModelInstanceStatusSummary() (*ServicesEc2ModelInstanceStatusSummary) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceStatusSummary")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceStatusSummary{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatusSummary) SetStatus2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatus()
func (jbobject *ServicesEc2ModelInstanceStatusSummary) GetStatus() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatus", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceStatusSummary withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatusSummary) WithStatus2(a string) *ServicesEc2ModelInstanceStatusSummary {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/InstanceStatusSummary", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceStatusSummary{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(com.amazonaws.services.ec2.model.SummaryStatus)
func (jbobject *ServicesEc2ModelInstanceStatusSummary) SetStatus(a ServicesEc2ModelSummaryStatusInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SummaryStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceStatusSummary withStatus(com.amazonaws.services.ec2.model.SummaryStatus)
func (jbobject *ServicesEc2ModelInstanceStatusSummary) WithStatus(a ServicesEc2ModelSummaryStatusInterface) *ServicesEc2ModelInstanceStatusSummary {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/InstanceStatusSummary", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SummaryStatus"))
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
	unique_x := &ServicesEc2ModelInstanceStatusSummary{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceStatusDetails> getDetails()
func (jbobject *ServicesEc2ModelInstanceStatusSummary) GetDetails() []*ServicesEc2ModelInstanceStatusDetails {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDetails", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInstanceStatusDetails)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setDetails(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStatusDetails>)
func (jbobject *ServicesEc2ModelInstanceStatusSummary) SetDetails(a []*ServicesEc2ModelInstanceStatusDetails)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDetails", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceStatusSummary withDetails(com.amazonaws.services.ec2.model.InstanceStatusDetails...)
func (jbobject *ServicesEc2ModelInstanceStatusSummary) WithDetails(a ...*ServicesEc2ModelInstanceStatusDetails) *ServicesEc2ModelInstanceStatusSummary {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceStatusDetails")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDetails", "com/amazonaws/services/ec2/model/InstanceStatusSummary", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceStatusDetails"))
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
	unique_x := &ServicesEc2ModelInstanceStatusSummary{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceStatusSummary withDetails(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStatusDetails>)
func (jbobject *ServicesEc2ModelInstanceStatusSummary) WithDetails2(a []*ServicesEc2ModelInstanceStatusDetails) *ServicesEc2ModelInstanceStatusSummary {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDetails", "com/amazonaws/services/ec2/model/InstanceStatusSummary", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelInstanceStatusSummary{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceStatusSummary) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceStatusSummary) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceStatusSummary) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceStatusSummary clone()
func (jbobject *ServicesEc2ModelInstanceStatusSummary) Clone() *ServicesEc2ModelInstanceStatusSummary {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceStatusSummary")
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
	unique_x := &ServicesEc2ModelInstanceStatusSummary{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceStatusSummary) Clone2() (*JavaLangObject, error) {
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


