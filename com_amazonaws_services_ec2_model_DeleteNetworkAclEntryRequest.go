package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDeleteNetworkAclEntryRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setNetworkAclId(java.lang.String)
	SetNetworkAclId(a string) 

	// public java.lang.String getNetworkAclId()
	GetNetworkAclId() string

	// public com.amazonaws.services.ec2.model.DeleteNetworkAclEntryRequest withNetworkAclId(java.lang.String)
	WithNetworkAclId(a string) *ServicesEc2ModelDeleteNetworkAclEntryRequest

	// public void setRuleNumber(java.lang.Integer)
	SetRuleNumber(a int) 

	// public java.lang.Integer getRuleNumber()
	GetRuleNumber() int

	// public com.amazonaws.services.ec2.model.DeleteNetworkAclEntryRequest withRuleNumber(java.lang.Integer)
	WithRuleNumber(a int) *ServicesEc2ModelDeleteNetworkAclEntryRequest

	// public void setEgress(java.lang.Boolean)
	SetEgress(a bool) 

	// public java.lang.Boolean getEgress()
	GetEgress() bool

	// public com.amazonaws.services.ec2.model.DeleteNetworkAclEntryRequest withEgress(java.lang.Boolean)
	WithEgress(a bool) *ServicesEc2ModelDeleteNetworkAclEntryRequest

	// public java.lang.Boolean isEgress()
	IsEgress() bool

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteNetworkAclEntryRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DeleteNetworkAclEntryRequest clone()
	Clone3() *ServicesEc2ModelDeleteNetworkAclEntryRequest
}

type ServicesEc2ModelDeleteNetworkAclEntryRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DeleteNetworkAclEntryRequest()
func NewServicesEc2ModelDeleteNetworkAclEntryRequest() (*ServicesEc2ModelDeleteNetworkAclEntryRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DeleteNetworkAclEntryRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDeleteNetworkAclEntryRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNetworkAclId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) SetNetworkAclId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkAclId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNetworkAclId()
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) GetNetworkAclId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkAclId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DeleteNetworkAclEntryRequest withNetworkAclId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) WithNetworkAclId(a string) *ServicesEc2ModelDeleteNetworkAclEntryRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkAclId", "com/amazonaws/services/ec2/model/DeleteNetworkAclEntryRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDeleteNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRuleNumber(java.lang.Integer)
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) SetRuleNumber(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRuleNumber", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getRuleNumber()
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) GetRuleNumber() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRuleNumber", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.DeleteNetworkAclEntryRequest withRuleNumber(java.lang.Integer)
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) WithRuleNumber(a int) *ServicesEc2ModelDeleteNetworkAclEntryRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRuleNumber", "com/amazonaws/services/ec2/model/DeleteNetworkAclEntryRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelDeleteNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEgress(java.lang.Boolean)
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) SetEgress(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEgress", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getEgress()
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) GetEgress() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEgress", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.DeleteNetworkAclEntryRequest withEgress(java.lang.Boolean)
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) WithEgress(a bool) *ServicesEc2ModelDeleteNetworkAclEntryRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEgress", "com/amazonaws/services/ec2/model/DeleteNetworkAclEntryRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelDeleteNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEgress()
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) IsEgress() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEgress", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteNetworkAclEntryRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) GetDryRunRequest() *Request {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDryRunRequest", "com/amazonaws/Request")
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
	unique_x := &Request{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DeleteNetworkAclEntryRequest clone()
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) Clone3() *ServicesEc2ModelDeleteNetworkAclEntryRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DeleteNetworkAclEntryRequest")
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
	unique_x := &ServicesEc2ModelDeleteNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) Clone() *AmazonWebServiceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/AmazonWebServiceRequest")
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
	unique_x := &AmazonWebServiceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDeleteNetworkAclEntryRequest) Clone2() (*JavaLangObject, error) {
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


