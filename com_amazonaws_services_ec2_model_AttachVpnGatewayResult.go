package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAttachVpnGatewayResultInterface interface {
	JavaLangObjectInterface

	// public void setVpcAttachment(com.amazonaws.services.ec2.model.VpcAttachment)
	SetVpcAttachment(a ServicesEc2ModelVpcAttachmentInterface) 

	// public com.amazonaws.services.ec2.model.VpcAttachment getVpcAttachment()
	GetVpcAttachment() *ServicesEc2ModelVpcAttachment

	// public com.amazonaws.services.ec2.model.AttachVpnGatewayResult withVpcAttachment(com.amazonaws.services.ec2.model.VpcAttachment)
	WithVpcAttachment(a ServicesEc2ModelVpcAttachmentInterface) *ServicesEc2ModelAttachVpnGatewayResult

	// public com.amazonaws.services.ec2.model.AttachVpnGatewayResult clone()
	Clone() *ServicesEc2ModelAttachVpnGatewayResult
}

type ServicesEc2ModelAttachVpnGatewayResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.AttachVpnGatewayResult()
func NewServicesEc2ModelAttachVpnGatewayResult() (*ServicesEc2ModelAttachVpnGatewayResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AttachVpnGatewayResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAttachVpnGatewayResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcAttachment(com.amazonaws.services.ec2.model.VpcAttachment)
func (jbobject *ServicesEc2ModelAttachVpnGatewayResult) SetVpcAttachment(a ServicesEc2ModelVpcAttachmentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcAttachment", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcAttachment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpcAttachment getVpcAttachment()
func (jbobject *ServicesEc2ModelAttachVpnGatewayResult) GetVpcAttachment() *ServicesEc2ModelVpcAttachment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcAttachment", "com/amazonaws/services/ec2/model/VpcAttachment")
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
	unique_x := &ServicesEc2ModelVpcAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AttachVpnGatewayResult withVpcAttachment(com.amazonaws.services.ec2.model.VpcAttachment)
func (jbobject *ServicesEc2ModelAttachVpnGatewayResult) WithVpcAttachment(a ServicesEc2ModelVpcAttachmentInterface) *ServicesEc2ModelAttachVpnGatewayResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcAttachment", "com/amazonaws/services/ec2/model/AttachVpnGatewayResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcAttachment"))
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
	unique_x := &ServicesEc2ModelAttachVpnGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAttachVpnGatewayResult) ToString() string {
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
func (jbobject *ServicesEc2ModelAttachVpnGatewayResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAttachVpnGatewayResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AttachVpnGatewayResult clone()
func (jbobject *ServicesEc2ModelAttachVpnGatewayResult) Clone() *ServicesEc2ModelAttachVpnGatewayResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AttachVpnGatewayResult")
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
	unique_x := &ServicesEc2ModelAttachVpnGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelAttachVpnGatewayResult) Clone2() (*JavaLangObject, error) {
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


