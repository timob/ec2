package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelReplaceNetworkAclAssociationResultInterface interface {
	JavaLangObjectInterface

	// public void setNewAssociationId(java.lang.String)
	SetNewAssociationId(a string) 

	// public java.lang.String getNewAssociationId()
	GetNewAssociationId() string

	// public com.amazonaws.services.ec2.model.ReplaceNetworkAclAssociationResult withNewAssociationId(java.lang.String)
	WithNewAssociationId(a string) *ServicesEc2ModelReplaceNetworkAclAssociationResult

	// public com.amazonaws.services.ec2.model.ReplaceNetworkAclAssociationResult clone()
	Clone() *ServicesEc2ModelReplaceNetworkAclAssociationResult
}

type ServicesEc2ModelReplaceNetworkAclAssociationResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ReplaceNetworkAclAssociationResult()
func NewServicesEc2ModelReplaceNetworkAclAssociationResult() (*ServicesEc2ModelReplaceNetworkAclAssociationResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ReplaceNetworkAclAssociationResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelReplaceNetworkAclAssociationResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNewAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelReplaceNetworkAclAssociationResult) SetNewAssociationId(a string)  {
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
func (jbobject *ServicesEc2ModelReplaceNetworkAclAssociationResult) GetNewAssociationId() string {
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

// public com.amazonaws.services.ec2.model.ReplaceNetworkAclAssociationResult withNewAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelReplaceNetworkAclAssociationResult) WithNewAssociationId(a string) *ServicesEc2ModelReplaceNetworkAclAssociationResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNewAssociationId", "com/amazonaws/services/ec2/model/ReplaceNetworkAclAssociationResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReplaceNetworkAclAssociationResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelReplaceNetworkAclAssociationResult) ToString() string {
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
func (jbobject *ServicesEc2ModelReplaceNetworkAclAssociationResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelReplaceNetworkAclAssociationResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ReplaceNetworkAclAssociationResult clone()
func (jbobject *ServicesEc2ModelReplaceNetworkAclAssociationResult) Clone() *ServicesEc2ModelReplaceNetworkAclAssociationResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ReplaceNetworkAclAssociationResult")
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
	unique_x := &ServicesEc2ModelReplaceNetworkAclAssociationResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelReplaceNetworkAclAssociationResult) Clone2() (*JavaLangObject, error) {
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


