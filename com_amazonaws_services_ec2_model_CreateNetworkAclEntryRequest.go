package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateNetworkAclEntryRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setNetworkAclId(java.lang.String)
	SetNetworkAclId(a string) 

	// public java.lang.String getNetworkAclId()
	GetNetworkAclId() string

	// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withNetworkAclId(java.lang.String)
	WithNetworkAclId(a string) *ServicesEc2ModelCreateNetworkAclEntryRequest

	// public void setRuleNumber(java.lang.Integer)
	SetRuleNumber(a int) 

	// public java.lang.Integer getRuleNumber()
	GetRuleNumber() int

	// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withRuleNumber(java.lang.Integer)
	WithRuleNumber(a int) *ServicesEc2ModelCreateNetworkAclEntryRequest

	// public void setProtocol(java.lang.String)
	SetProtocol(a string) 

	// public java.lang.String getProtocol()
	GetProtocol() string

	// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withProtocol(java.lang.String)
	WithProtocol(a string) *ServicesEc2ModelCreateNetworkAclEntryRequest

	// public void setRuleAction(java.lang.String)
	SetRuleAction2(a string) 

	// public java.lang.String getRuleAction()
	GetRuleAction() string

	// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withRuleAction(java.lang.String)
	WithRuleAction2(a string) *ServicesEc2ModelCreateNetworkAclEntryRequest

	// public void setRuleAction(com.amazonaws.services.ec2.model.RuleAction)
	SetRuleAction(a ServicesEc2ModelRuleActionInterface) 

	// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withRuleAction(com.amazonaws.services.ec2.model.RuleAction)
	WithRuleAction(a ServicesEc2ModelRuleActionInterface) *ServicesEc2ModelCreateNetworkAclEntryRequest

	// public void setEgress(java.lang.Boolean)
	SetEgress(a bool) 

	// public java.lang.Boolean getEgress()
	GetEgress() bool

	// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withEgress(java.lang.Boolean)
	WithEgress(a bool) *ServicesEc2ModelCreateNetworkAclEntryRequest

	// public java.lang.Boolean isEgress()
	IsEgress() bool

	// public void setCidrBlock(java.lang.String)
	SetCidrBlock(a string) 

	// public java.lang.String getCidrBlock()
	GetCidrBlock() string

	// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withCidrBlock(java.lang.String)
	WithCidrBlock(a string) *ServicesEc2ModelCreateNetworkAclEntryRequest

	// public void setIcmpTypeCode(com.amazonaws.services.ec2.model.IcmpTypeCode)
	SetIcmpTypeCode(a ServicesEc2ModelIcmpTypeCodeInterface) 

	// public com.amazonaws.services.ec2.model.IcmpTypeCode getIcmpTypeCode()
	GetIcmpTypeCode() *ServicesEc2ModelIcmpTypeCode

	// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withIcmpTypeCode(com.amazonaws.services.ec2.model.IcmpTypeCode)
	WithIcmpTypeCode(a ServicesEc2ModelIcmpTypeCodeInterface) *ServicesEc2ModelCreateNetworkAclEntryRequest

	// public void setPortRange(com.amazonaws.services.ec2.model.PortRange)
	SetPortRange(a ServicesEc2ModelPortRangeInterface) 

	// public com.amazonaws.services.ec2.model.PortRange getPortRange()
	GetPortRange() *ServicesEc2ModelPortRange

	// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withPortRange(com.amazonaws.services.ec2.model.PortRange)
	WithPortRange(a ServicesEc2ModelPortRangeInterface) *ServicesEc2ModelCreateNetworkAclEntryRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest clone()
	Clone3() *ServicesEc2ModelCreateNetworkAclEntryRequest
}

type ServicesEc2ModelCreateNetworkAclEntryRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest()
func NewServicesEc2ModelCreateNetworkAclEntryRequest() (*ServicesEc2ModelCreateNetworkAclEntryRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateNetworkAclEntryRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateNetworkAclEntryRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNetworkAclId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) SetNetworkAclId(a string)  {
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
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) GetNetworkAclId() string {
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

// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withNetworkAclId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) WithNetworkAclId(a string) *ServicesEc2ModelCreateNetworkAclEntryRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkAclId", "com/amazonaws/services/ec2/model/CreateNetworkAclEntryRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRuleNumber(java.lang.Integer)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) SetRuleNumber(a int)  {
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
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) GetRuleNumber() int {
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

// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withRuleNumber(java.lang.Integer)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) WithRuleNumber(a int) *ServicesEc2ModelCreateNetworkAclEntryRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRuleNumber", "com/amazonaws/services/ec2/model/CreateNetworkAclEntryRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelCreateNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProtocol(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) SetProtocol(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProtocol", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getProtocol()
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) GetProtocol() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProtocol", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withProtocol(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) WithProtocol(a string) *ServicesEc2ModelCreateNetworkAclEntryRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProtocol", "com/amazonaws/services/ec2/model/CreateNetworkAclEntryRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRuleAction(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) SetRuleAction2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRuleAction", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRuleAction()
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) GetRuleAction() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRuleAction", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withRuleAction(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) WithRuleAction2(a string) *ServicesEc2ModelCreateNetworkAclEntryRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRuleAction", "com/amazonaws/services/ec2/model/CreateNetworkAclEntryRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRuleAction(com.amazonaws.services.ec2.model.RuleAction)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) SetRuleAction(a ServicesEc2ModelRuleActionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRuleAction", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/RuleAction"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withRuleAction(com.amazonaws.services.ec2.model.RuleAction)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) WithRuleAction(a ServicesEc2ModelRuleActionInterface) *ServicesEc2ModelCreateNetworkAclEntryRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRuleAction", "com/amazonaws/services/ec2/model/CreateNetworkAclEntryRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RuleAction"))
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
	unique_x := &ServicesEc2ModelCreateNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEgress(java.lang.Boolean)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) SetEgress(a bool)  {
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
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) GetEgress() bool {
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

// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withEgress(java.lang.Boolean)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) WithEgress(a bool) *ServicesEc2ModelCreateNetworkAclEntryRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEgress", "com/amazonaws/services/ec2/model/CreateNetworkAclEntryRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelCreateNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEgress()
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) IsEgress() bool {
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

// public void setCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) SetCidrBlock(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCidrBlock", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCidrBlock()
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) GetCidrBlock() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCidrBlock", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) WithCidrBlock(a string) *ServicesEc2ModelCreateNetworkAclEntryRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCidrBlock", "com/amazonaws/services/ec2/model/CreateNetworkAclEntryRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIcmpTypeCode(com.amazonaws.services.ec2.model.IcmpTypeCode)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) SetIcmpTypeCode(a ServicesEc2ModelIcmpTypeCodeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIcmpTypeCode", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/IcmpTypeCode"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.IcmpTypeCode getIcmpTypeCode()
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) GetIcmpTypeCode() *ServicesEc2ModelIcmpTypeCode {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIcmpTypeCode", "com/amazonaws/services/ec2/model/IcmpTypeCode")
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
	unique_x := &ServicesEc2ModelIcmpTypeCode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withIcmpTypeCode(com.amazonaws.services.ec2.model.IcmpTypeCode)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) WithIcmpTypeCode(a ServicesEc2ModelIcmpTypeCodeInterface) *ServicesEc2ModelCreateNetworkAclEntryRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIcmpTypeCode", "com/amazonaws/services/ec2/model/CreateNetworkAclEntryRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/IcmpTypeCode"))
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
	unique_x := &ServicesEc2ModelCreateNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPortRange(com.amazonaws.services.ec2.model.PortRange)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) SetPortRange(a ServicesEc2ModelPortRangeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPortRange", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/PortRange"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.PortRange getPortRange()
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) GetPortRange() *ServicesEc2ModelPortRange {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPortRange", "com/amazonaws/services/ec2/model/PortRange")
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
	unique_x := &ServicesEc2ModelPortRange{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest withPortRange(com.amazonaws.services.ec2.model.PortRange)
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) WithPortRange(a ServicesEc2ModelPortRangeInterface) *ServicesEc2ModelCreateNetworkAclEntryRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPortRange", "com/amazonaws/services/ec2/model/CreateNetworkAclEntryRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PortRange"))
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
	unique_x := &ServicesEc2ModelCreateNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest clone()
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) Clone3() *ServicesEc2ModelCreateNetworkAclEntryRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateNetworkAclEntryRequest")
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
	unique_x := &ServicesEc2ModelCreateNetworkAclEntryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreateNetworkAclEntryRequest) Clone2() (*JavaLangObject, error) {
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


