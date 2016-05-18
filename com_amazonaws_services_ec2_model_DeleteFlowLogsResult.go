package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDeleteFlowLogsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.UnsuccessfulItem> getUnsuccessful()
	GetUnsuccessful() []*ServicesEc2ModelUnsuccessfulItem

	// public void setUnsuccessful(java.util.Collection<com.amazonaws.services.ec2.model.UnsuccessfulItem>)
	SetUnsuccessful(a []*ServicesEc2ModelUnsuccessfulItem) 

	// public com.amazonaws.services.ec2.model.DeleteFlowLogsResult withUnsuccessful(com.amazonaws.services.ec2.model.UnsuccessfulItem...)
	WithUnsuccessful(a ...*ServicesEc2ModelUnsuccessfulItem) *ServicesEc2ModelDeleteFlowLogsResult

	// public com.amazonaws.services.ec2.model.DeleteFlowLogsResult withUnsuccessful(java.util.Collection<com.amazonaws.services.ec2.model.UnsuccessfulItem>)
	WithUnsuccessful2(a []*ServicesEc2ModelUnsuccessfulItem) *ServicesEc2ModelDeleteFlowLogsResult

	// public com.amazonaws.services.ec2.model.DeleteFlowLogsResult clone()
	Clone() *ServicesEc2ModelDeleteFlowLogsResult
}

type ServicesEc2ModelDeleteFlowLogsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DeleteFlowLogsResult()
func NewServicesEc2ModelDeleteFlowLogsResult() (*ServicesEc2ModelDeleteFlowLogsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DeleteFlowLogsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDeleteFlowLogsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.UnsuccessfulItem> getUnsuccessful()
func (jbobject *ServicesEc2ModelDeleteFlowLogsResult) GetUnsuccessful() []*ServicesEc2ModelUnsuccessfulItem {
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
func (jbobject *ServicesEc2ModelDeleteFlowLogsResult) SetUnsuccessful(a []*ServicesEc2ModelUnsuccessfulItem)  {
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

// public com.amazonaws.services.ec2.model.DeleteFlowLogsResult withUnsuccessful(com.amazonaws.services.ec2.model.UnsuccessfulItem...)
func (jbobject *ServicesEc2ModelDeleteFlowLogsResult) WithUnsuccessful(a ...*ServicesEc2ModelUnsuccessfulItem) *ServicesEc2ModelDeleteFlowLogsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/UnsuccessfulItem")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUnsuccessful", "com/amazonaws/services/ec2/model/DeleteFlowLogsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/UnsuccessfulItem"))
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
	unique_x := &ServicesEc2ModelDeleteFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DeleteFlowLogsResult withUnsuccessful(java.util.Collection<com.amazonaws.services.ec2.model.UnsuccessfulItem>)
func (jbobject *ServicesEc2ModelDeleteFlowLogsResult) WithUnsuccessful2(a []*ServicesEc2ModelUnsuccessfulItem) *ServicesEc2ModelDeleteFlowLogsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUnsuccessful", "com/amazonaws/services/ec2/model/DeleteFlowLogsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDeleteFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDeleteFlowLogsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDeleteFlowLogsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDeleteFlowLogsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DeleteFlowLogsResult clone()
func (jbobject *ServicesEc2ModelDeleteFlowLogsResult) Clone() *ServicesEc2ModelDeleteFlowLogsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DeleteFlowLogsResult")
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
	unique_x := &ServicesEc2ModelDeleteFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDeleteFlowLogsResult) Clone2() (*JavaLangObject, error) {
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


