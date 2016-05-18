package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelSpotFleetRequestsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem> getUnsuccessfulFleetRequests()
	GetUnsuccessfulFleetRequests() []*ServicesEc2ModelCancelSpotFleetRequestsErrorItem

	// public void setUnsuccessfulFleetRequests(java.util.Collection<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem>)
	SetUnsuccessfulFleetRequests(a []*ServicesEc2ModelCancelSpotFleetRequestsErrorItem) 

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsResult withUnsuccessfulFleetRequests(com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem...)
	WithUnsuccessfulFleetRequests(a ...*ServicesEc2ModelCancelSpotFleetRequestsErrorItem) *ServicesEc2ModelCancelSpotFleetRequestsResult

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsResult withUnsuccessfulFleetRequests(java.util.Collection<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem>)
	WithUnsuccessfulFleetRequests2(a []*ServicesEc2ModelCancelSpotFleetRequestsErrorItem) *ServicesEc2ModelCancelSpotFleetRequestsResult

	// public java.util.List<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem> getSuccessfulFleetRequests()
	GetSuccessfulFleetRequests() []*ServicesEc2ModelCancelSpotFleetRequestsSuccessItem

	// public void setSuccessfulFleetRequests(java.util.Collection<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem>)
	SetSuccessfulFleetRequests(a []*ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) 

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsResult withSuccessfulFleetRequests(com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem...)
	WithSuccessfulFleetRequests(a ...*ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) *ServicesEc2ModelCancelSpotFleetRequestsResult

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsResult withSuccessfulFleetRequests(java.util.Collection<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem>)
	WithSuccessfulFleetRequests2(a []*ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) *ServicesEc2ModelCancelSpotFleetRequestsResult

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsResult clone()
	Clone() *ServicesEc2ModelCancelSpotFleetRequestsResult
}

type ServicesEc2ModelCancelSpotFleetRequestsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsResult()
func NewServicesEc2ModelCancelSpotFleetRequestsResult() (*ServicesEc2ModelCancelSpotFleetRequestsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelSpotFleetRequestsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelSpotFleetRequestsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem> getUnsuccessfulFleetRequests()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsResult) GetUnsuccessfulFleetRequests() []*ServicesEc2ModelCancelSpotFleetRequestsErrorItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUnsuccessfulFleetRequests", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelCancelSpotFleetRequestsErrorItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setUnsuccessfulFleetRequests(java.util.Collection<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem>)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsResult) SetUnsuccessfulFleetRequests(a []*ServicesEc2ModelCancelSpotFleetRequestsErrorItem)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUnsuccessfulFleetRequests", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsResult withUnsuccessfulFleetRequests(com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem...)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsResult) WithUnsuccessfulFleetRequests(a ...*ServicesEc2ModelCancelSpotFleetRequestsErrorItem) *ServicesEc2ModelCancelSpotFleetRequestsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsErrorItem")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUnsuccessfulFleetRequests", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelSpotFleetRequestsErrorItem"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsResult withUnsuccessfulFleetRequests(java.util.Collection<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem>)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsResult) WithUnsuccessfulFleetRequests2(a []*ServicesEc2ModelCancelSpotFleetRequestsErrorItem) *ServicesEc2ModelCancelSpotFleetRequestsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUnsuccessfulFleetRequests", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem> getSuccessfulFleetRequests()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsResult) GetSuccessfulFleetRequests() []*ServicesEc2ModelCancelSpotFleetRequestsSuccessItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSuccessfulFleetRequests", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelCancelSpotFleetRequestsSuccessItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSuccessfulFleetRequests(java.util.Collection<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem>)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsResult) SetSuccessfulFleetRequests(a []*ServicesEc2ModelCancelSpotFleetRequestsSuccessItem)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSuccessfulFleetRequests", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsResult withSuccessfulFleetRequests(com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem...)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsResult) WithSuccessfulFleetRequests(a ...*ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) *ServicesEc2ModelCancelSpotFleetRequestsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsSuccessItem")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSuccessfulFleetRequests", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelSpotFleetRequestsSuccessItem"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsResult withSuccessfulFleetRequests(java.util.Collection<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem>)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsResult) WithSuccessfulFleetRequests2(a []*ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) *ServicesEc2ModelCancelSpotFleetRequestsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSuccessfulFleetRequests", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsResult clone()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsResult) Clone() *ServicesEc2ModelCancelSpotFleetRequestsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsResult")
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsResult) Clone2() (*JavaLangObject, error) {
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


