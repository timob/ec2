package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPurchaseScheduledInstancesRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesRequest withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelPurchaseScheduledInstancesRequest

	// public java.util.List<com.amazonaws.services.ec2.model.PurchaseRequest> getPurchaseRequests()
	GetPurchaseRequests() []*ServicesEc2ModelPurchaseRequest

	// public void setPurchaseRequests(java.util.Collection<com.amazonaws.services.ec2.model.PurchaseRequest>)
	SetPurchaseRequests(a []*ServicesEc2ModelPurchaseRequest) 

	// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesRequest withPurchaseRequests(com.amazonaws.services.ec2.model.PurchaseRequest...)
	WithPurchaseRequests(a ...*ServicesEc2ModelPurchaseRequest) *ServicesEc2ModelPurchaseScheduledInstancesRequest

	// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesRequest withPurchaseRequests(java.util.Collection<com.amazonaws.services.ec2.model.PurchaseRequest>)
	WithPurchaseRequests2(a []*ServicesEc2ModelPurchaseRequest) *ServicesEc2ModelPurchaseScheduledInstancesRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.PurchaseScheduledInstancesRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesRequest clone()
	Clone3() *ServicesEc2ModelPurchaseScheduledInstancesRequest
}

type ServicesEc2ModelPurchaseScheduledInstancesRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesRequest()
func NewServicesEc2ModelPurchaseScheduledInstancesRequest() (*ServicesEc2ModelPurchaseScheduledInstancesRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PurchaseScheduledInstancesRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPurchaseScheduledInstancesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) SetClientToken(a string)  {
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
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) GetClientToken() string {
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

// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesRequest withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) WithClientToken(a string) *ServicesEc2ModelPurchaseScheduledInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/PurchaseScheduledInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPurchaseScheduledInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.PurchaseRequest> getPurchaseRequests()
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) GetPurchaseRequests() []*ServicesEc2ModelPurchaseRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPurchaseRequests", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelPurchaseRequest)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setPurchaseRequests(java.util.Collection<com.amazonaws.services.ec2.model.PurchaseRequest>)
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) SetPurchaseRequests(a []*ServicesEc2ModelPurchaseRequest)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPurchaseRequests", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesRequest withPurchaseRequests(com.amazonaws.services.ec2.model.PurchaseRequest...)
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) WithPurchaseRequests(a ...*ServicesEc2ModelPurchaseRequest) *ServicesEc2ModelPurchaseScheduledInstancesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/PurchaseRequest")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPurchaseRequests", "com/amazonaws/services/ec2/model/PurchaseScheduledInstancesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PurchaseRequest"))
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
	unique_x := &ServicesEc2ModelPurchaseScheduledInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesRequest withPurchaseRequests(java.util.Collection<com.amazonaws.services.ec2.model.PurchaseRequest>)
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) WithPurchaseRequests2(a []*ServicesEc2ModelPurchaseRequest) *ServicesEc2ModelPurchaseScheduledInstancesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPurchaseRequests", "com/amazonaws/services/ec2/model/PurchaseScheduledInstancesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelPurchaseScheduledInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.PurchaseScheduledInstancesRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesRequest clone()
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) Clone3() *ServicesEc2ModelPurchaseScheduledInstancesRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/PurchaseScheduledInstancesRequest")
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
	unique_x := &ServicesEc2ModelPurchaseScheduledInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesRequest) Clone2() (*JavaLangObject, error) {
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


