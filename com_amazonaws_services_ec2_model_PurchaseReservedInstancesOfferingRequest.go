package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPurchaseReservedInstancesOfferingRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setReservedInstancesOfferingId(java.lang.String)
	SetReservedInstancesOfferingId(a string) 

	// public java.lang.String getReservedInstancesOfferingId()
	GetReservedInstancesOfferingId() string

	// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest withReservedInstancesOfferingId(java.lang.String)
	WithReservedInstancesOfferingId(a string) *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest

	// public void setInstanceCount(java.lang.Integer)
	SetInstanceCount(a int) 

	// public java.lang.Integer getInstanceCount()
	GetInstanceCount() int

	// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest withInstanceCount(java.lang.Integer)
	WithInstanceCount(a int) *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest

	// public void setLimitPrice(com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice)
	SetLimitPrice(a ServicesEc2ModelReservedInstanceLimitPriceInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice getLimitPrice()
	GetLimitPrice() *ServicesEc2ModelReservedInstanceLimitPrice

	// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest withLimitPrice(com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice)
	WithLimitPrice(a ServicesEc2ModelReservedInstanceLimitPriceInterface) *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest clone()
	Clone3() *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest
}

type ServicesEc2ModelPurchaseReservedInstancesOfferingRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest()
func NewServicesEc2ModelPurchaseReservedInstancesOfferingRequest() (*ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PurchaseReservedInstancesOfferingRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPurchaseReservedInstancesOfferingRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest(java.lang.String, java.lang.Integer)
func NewServicesEc2ModelPurchaseReservedInstancesOfferingRequest2(a string, b int) (*ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PurchaseReservedInstancesOfferingRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelPurchaseReservedInstancesOfferingRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setReservedInstancesOfferingId(java.lang.String)
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) SetReservedInstancesOfferingId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesOfferingId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getReservedInstancesOfferingId()
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) GetReservedInstancesOfferingId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesOfferingId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest withReservedInstancesOfferingId(java.lang.String)
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) WithReservedInstancesOfferingId(a string) *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesOfferingId", "com/amazonaws/services/ec2/model/PurchaseReservedInstancesOfferingRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPurchaseReservedInstancesOfferingRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) SetInstanceCount(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceCount", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getInstanceCount()
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) GetInstanceCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceCount", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest withInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) WithInstanceCount(a int) *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceCount", "com/amazonaws/services/ec2/model/PurchaseReservedInstancesOfferingRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelPurchaseReservedInstancesOfferingRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLimitPrice(com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice)
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) SetLimitPrice(a ServicesEc2ModelReservedInstanceLimitPriceInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLimitPrice", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstanceLimitPrice"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice getLimitPrice()
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) GetLimitPrice() *ServicesEc2ModelReservedInstanceLimitPrice {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLimitPrice", "com/amazonaws/services/ec2/model/ReservedInstanceLimitPrice")
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
	unique_x := &ServicesEc2ModelReservedInstanceLimitPrice{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest withLimitPrice(com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice)
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) WithLimitPrice(a ServicesEc2ModelReservedInstanceLimitPriceInterface) *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLimitPrice", "com/amazonaws/services/ec2/model/PurchaseReservedInstancesOfferingRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstanceLimitPrice"))
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
	unique_x := &ServicesEc2ModelPurchaseReservedInstancesOfferingRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest clone()
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) Clone3() *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/PurchaseReservedInstancesOfferingRequest")
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
	unique_x := &ServicesEc2ModelPurchaseReservedInstancesOfferingRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingRequest) Clone2() (*JavaLangObject, error) {
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


