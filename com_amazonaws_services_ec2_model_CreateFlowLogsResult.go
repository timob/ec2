package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateFlowLogsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<java.lang.String> getFlowLogIds()
	GetFlowLogIds() []string

	// public void setFlowLogIds(java.util.Collection<java.lang.String>)
	SetFlowLogIds(a []string) 

	// public com.amazonaws.services.ec2.model.CreateFlowLogsResult withFlowLogIds(java.lang.String...)
	WithFlowLogIds(a ...string) *ServicesEc2ModelCreateFlowLogsResult

	// public com.amazonaws.services.ec2.model.CreateFlowLogsResult withFlowLogIds(java.util.Collection<java.lang.String>)
	WithFlowLogIds2(a []string) *ServicesEc2ModelCreateFlowLogsResult

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.CreateFlowLogsResult withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelCreateFlowLogsResult

	// public java.util.List<com.amazonaws.services.ec2.model.UnsuccessfulItem> getUnsuccessful()
	GetUnsuccessful() []*ServicesEc2ModelUnsuccessfulItem

	// public void setUnsuccessful(java.util.Collection<com.amazonaws.services.ec2.model.UnsuccessfulItem>)
	SetUnsuccessful(a []*ServicesEc2ModelUnsuccessfulItem) 

	// public com.amazonaws.services.ec2.model.CreateFlowLogsResult withUnsuccessful(com.amazonaws.services.ec2.model.UnsuccessfulItem...)
	WithUnsuccessful(a ...*ServicesEc2ModelUnsuccessfulItem) *ServicesEc2ModelCreateFlowLogsResult

	// public com.amazonaws.services.ec2.model.CreateFlowLogsResult withUnsuccessful(java.util.Collection<com.amazonaws.services.ec2.model.UnsuccessfulItem>)
	WithUnsuccessful2(a []*ServicesEc2ModelUnsuccessfulItem) *ServicesEc2ModelCreateFlowLogsResult

	// public com.amazonaws.services.ec2.model.CreateFlowLogsResult clone()
	Clone() *ServicesEc2ModelCreateFlowLogsResult
}

type ServicesEc2ModelCreateFlowLogsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateFlowLogsResult()
func NewServicesEc2ModelCreateFlowLogsResult() (*ServicesEc2ModelCreateFlowLogsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateFlowLogsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateFlowLogsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getFlowLogIds()
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) GetFlowLogIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFlowLogIds", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setFlowLogIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) SetFlowLogIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFlowLogIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateFlowLogsResult withFlowLogIds(java.lang.String...)
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) WithFlowLogIds(a ...string) *ServicesEc2ModelCreateFlowLogsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFlowLogIds", "com/amazonaws/services/ec2/model/CreateFlowLogsResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateFlowLogsResult withFlowLogIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) WithFlowLogIds2(a []string) *ServicesEc2ModelCreateFlowLogsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFlowLogIds", "com/amazonaws/services/ec2/model/CreateFlowLogsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) SetClientToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClientToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getClientToken()
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) GetClientToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateFlowLogsResult withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) WithClientToken(a string) *ServicesEc2ModelCreateFlowLogsResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/CreateFlowLogsResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.UnsuccessfulItem> getUnsuccessful()
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) GetUnsuccessful() []*ServicesEc2ModelUnsuccessfulItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUnsuccessful", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelUnsuccessfulItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setUnsuccessful(java.util.Collection<com.amazonaws.services.ec2.model.UnsuccessfulItem>)
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) SetUnsuccessful(a []*ServicesEc2ModelUnsuccessfulItem)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUnsuccessful", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateFlowLogsResult withUnsuccessful(com.amazonaws.services.ec2.model.UnsuccessfulItem...)
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) WithUnsuccessful(a ...*ServicesEc2ModelUnsuccessfulItem) *ServicesEc2ModelCreateFlowLogsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/UnsuccessfulItem")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUnsuccessful", "com/amazonaws/services/ec2/model/CreateFlowLogsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/UnsuccessfulItem"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateFlowLogsResult withUnsuccessful(java.util.Collection<com.amazonaws.services.ec2.model.UnsuccessfulItem>)
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) WithUnsuccessful2(a []*ServicesEc2ModelUnsuccessfulItem) *ServicesEc2ModelCreateFlowLogsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUnsuccessful", "com/amazonaws/services/ec2/model/CreateFlowLogsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateFlowLogsResult clone()
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) Clone() *ServicesEc2ModelCreateFlowLogsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateFlowLogsResult")
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
	unique_x := &ServicesEc2ModelCreateFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateFlowLogsResult) Clone2() (*JavaLangObject, error) {
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


