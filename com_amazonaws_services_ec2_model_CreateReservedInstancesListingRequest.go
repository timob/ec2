package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateReservedInstancesListingRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setReservedInstancesId(java.lang.String)
	SetReservedInstancesId(a string) 

	// public java.lang.String getReservedInstancesId()
	GetReservedInstancesId() string

	// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest withReservedInstancesId(java.lang.String)
	WithReservedInstancesId(a string) *ServicesEc2ModelCreateReservedInstancesListingRequest

	// public void setInstanceCount(java.lang.Integer)
	SetInstanceCount(a int) 

	// public java.lang.Integer getInstanceCount()
	GetInstanceCount() int

	// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest withInstanceCount(java.lang.Integer)
	WithInstanceCount(a int) *ServicesEc2ModelCreateReservedInstancesListingRequest

	// public java.util.List<com.amazonaws.services.ec2.model.PriceScheduleSpecification> getPriceSchedules()
	GetPriceSchedules() []*ServicesEc2ModelPriceScheduleSpecification

	// public void setPriceSchedules(java.util.Collection<com.amazonaws.services.ec2.model.PriceScheduleSpecification>)
	SetPriceSchedules(a []*ServicesEc2ModelPriceScheduleSpecification) 

	// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest withPriceSchedules(com.amazonaws.services.ec2.model.PriceScheduleSpecification...)
	WithPriceSchedules(a ...*ServicesEc2ModelPriceScheduleSpecification) *ServicesEc2ModelCreateReservedInstancesListingRequest

	// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest withPriceSchedules(java.util.Collection<com.amazonaws.services.ec2.model.PriceScheduleSpecification>)
	WithPriceSchedules2(a []*ServicesEc2ModelPriceScheduleSpecification) *ServicesEc2ModelCreateReservedInstancesListingRequest

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelCreateReservedInstancesListingRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest clone()
	Clone3() *ServicesEc2ModelCreateReservedInstancesListingRequest
}

type ServicesEc2ModelCreateReservedInstancesListingRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest()
func NewServicesEc2ModelCreateReservedInstancesListingRequest() (*ServicesEc2ModelCreateReservedInstancesListingRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateReservedInstancesListingRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateReservedInstancesListingRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setReservedInstancesId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) SetReservedInstancesId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getReservedInstancesId()
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) GetReservedInstancesId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest withReservedInstancesId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) WithReservedInstancesId(a string) *ServicesEc2ModelCreateReservedInstancesListingRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesId", "com/amazonaws/services/ec2/model/CreateReservedInstancesListingRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateReservedInstancesListingRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) SetInstanceCount(a int)  {
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
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) GetInstanceCount() int {
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

// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest withInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) WithInstanceCount(a int) *ServicesEc2ModelCreateReservedInstancesListingRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceCount", "com/amazonaws/services/ec2/model/CreateReservedInstancesListingRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelCreateReservedInstancesListingRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.PriceScheduleSpecification> getPriceSchedules()
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) GetPriceSchedules() []*ServicesEc2ModelPriceScheduleSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPriceSchedules", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelPriceScheduleSpecification)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setPriceSchedules(java.util.Collection<com.amazonaws.services.ec2.model.PriceScheduleSpecification>)
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) SetPriceSchedules(a []*ServicesEc2ModelPriceScheduleSpecification)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPriceSchedules", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest withPriceSchedules(com.amazonaws.services.ec2.model.PriceScheduleSpecification...)
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) WithPriceSchedules(a ...*ServicesEc2ModelPriceScheduleSpecification) *ServicesEc2ModelCreateReservedInstancesListingRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/PriceScheduleSpecification")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPriceSchedules", "com/amazonaws/services/ec2/model/CreateReservedInstancesListingRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PriceScheduleSpecification"))
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
	unique_x := &ServicesEc2ModelCreateReservedInstancesListingRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest withPriceSchedules(java.util.Collection<com.amazonaws.services.ec2.model.PriceScheduleSpecification>)
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) WithPriceSchedules2(a []*ServicesEc2ModelPriceScheduleSpecification) *ServicesEc2ModelCreateReservedInstancesListingRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPriceSchedules", "com/amazonaws/services/ec2/model/CreateReservedInstancesListingRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCreateReservedInstancesListingRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) SetClientToken(a string)  {
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
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) GetClientToken() string {
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

// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) WithClientToken(a string) *ServicesEc2ModelCreateReservedInstancesListingRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/CreateReservedInstancesListingRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateReservedInstancesListingRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest clone()
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) Clone3() *ServicesEc2ModelCreateReservedInstancesListingRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateReservedInstancesListingRequest")
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
	unique_x := &ServicesEc2ModelCreateReservedInstancesListingRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingRequest) Clone2() (*JavaLangObject, error) {
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


