package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelReplaceRouteTableAssociationResultInterface interface {
	JavaLangObjectInterface

	// public void setNewAssociationId(java.lang.String)
	SetNewAssociationId(a string) 

	// public java.lang.String getNewAssociationId()
	GetNewAssociationId() string

	// public com.amazonaws.services.ec2.model.ReplaceRouteTableAssociationResult withNewAssociationId(java.lang.String)
	WithNewAssociationId(a string) *ServicesEc2ModelReplaceRouteTableAssociationResult

	// public com.amazonaws.services.ec2.model.ReplaceRouteTableAssociationResult clone()
	Clone() *ServicesEc2ModelReplaceRouteTableAssociationResult
}

type ServicesEc2ModelReplaceRouteTableAssociationResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ReplaceRouteTableAssociationResult()
func NewServicesEc2ModelReplaceRouteTableAssociationResult() (*ServicesEc2ModelReplaceRouteTableAssociationResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ReplaceRouteTableAssociationResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelReplaceRouteTableAssociationResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNewAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelReplaceRouteTableAssociationResult) SetNewAssociationId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNewAssociationId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNewAssociationId()
func (jbobject *ServicesEc2ModelReplaceRouteTableAssociationResult) GetNewAssociationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNewAssociationId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReplaceRouteTableAssociationResult withNewAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelReplaceRouteTableAssociationResult) WithNewAssociationId(a string) *ServicesEc2ModelReplaceRouteTableAssociationResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNewAssociationId", "com/amazonaws/services/ec2/model/ReplaceRouteTableAssociationResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReplaceRouteTableAssociationResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelReplaceRouteTableAssociationResult) ToString() string {
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
func (jbobject *ServicesEc2ModelReplaceRouteTableAssociationResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelReplaceRouteTableAssociationResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ReplaceRouteTableAssociationResult clone()
func (jbobject *ServicesEc2ModelReplaceRouteTableAssociationResult) Clone() *ServicesEc2ModelReplaceRouteTableAssociationResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ReplaceRouteTableAssociationResult")
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
	unique_x := &ServicesEc2ModelReplaceRouteTableAssociationResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelReplaceRouteTableAssociationResult) Clone2() (*JavaLangObject, error) {
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


