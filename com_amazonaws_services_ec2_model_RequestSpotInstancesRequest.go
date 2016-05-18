package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelRequestSpotInstancesRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setSpotPrice(java.lang.String)
	SetSpotPrice(a string) 

	// public java.lang.String getSpotPrice()
	GetSpotPrice() string

	// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withSpotPrice(java.lang.String)
	WithSpotPrice(a string) *ServicesEc2ModelRequestSpotInstancesRequest

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelRequestSpotInstancesRequest

	// public void setInstanceCount(java.lang.Integer)
	SetInstanceCount(a int) 

	// public java.lang.Integer getInstanceCount()
	GetInstanceCount() int

	// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withInstanceCount(java.lang.Integer)
	WithInstanceCount(a int) *ServicesEc2ModelRequestSpotInstancesRequest

	// public void setType(java.lang.String)
	SetType2(a string) 

	// public java.lang.String getType()
	GetType() string

	// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withType(java.lang.String)
	WithType2(a string) *ServicesEc2ModelRequestSpotInstancesRequest

	// public void setType(com.amazonaws.services.ec2.model.SpotInstanceType)
	SetType(a ServicesEc2ModelSpotInstanceTypeInterface) 

	// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withType(com.amazonaws.services.ec2.model.SpotInstanceType)
	WithType(a ServicesEc2ModelSpotInstanceTypeInterface) *ServicesEc2ModelRequestSpotInstancesRequest

	// public void setValidFrom(java.util.Date)
	SetValidFrom(a time.Time) 

	// public java.util.Date getValidFrom()
	GetValidFrom() time.Time

	// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withValidFrom(java.util.Date)
	WithValidFrom(a time.Time) *ServicesEc2ModelRequestSpotInstancesRequest

	// public void setValidUntil(java.util.Date)
	SetValidUntil(a time.Time) 

	// public java.util.Date getValidUntil()
	GetValidUntil() time.Time

	// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withValidUntil(java.util.Date)
	WithValidUntil(a time.Time) *ServicesEc2ModelRequestSpotInstancesRequest

	// public void setLaunchGroup(java.lang.String)
	SetLaunchGroup(a string) 

	// public java.lang.String getLaunchGroup()
	GetLaunchGroup() string

	// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withLaunchGroup(java.lang.String)
	WithLaunchGroup(a string) *ServicesEc2ModelRequestSpotInstancesRequest

	// public void setAvailabilityZoneGroup(java.lang.String)
	SetAvailabilityZoneGroup(a string) 

	// public java.lang.String getAvailabilityZoneGroup()
	GetAvailabilityZoneGroup() string

	// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withAvailabilityZoneGroup(java.lang.String)
	WithAvailabilityZoneGroup(a string) *ServicesEc2ModelRequestSpotInstancesRequest

	// public void setBlockDurationMinutes(java.lang.Integer)
	SetBlockDurationMinutes(a int) 

	// public java.lang.Integer getBlockDurationMinutes()
	GetBlockDurationMinutes() int

	// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withBlockDurationMinutes(java.lang.Integer)
	WithBlockDurationMinutes(a int) *ServicesEc2ModelRequestSpotInstancesRequest

	// public void setLaunchSpecification(com.amazonaws.services.ec2.model.LaunchSpecification)
	SetLaunchSpecification(a ServicesEc2ModelLaunchSpecificationInterface) 

	// public com.amazonaws.services.ec2.model.LaunchSpecification getLaunchSpecification()
	GetLaunchSpecification() *ServicesEc2ModelLaunchSpecification

	// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withLaunchSpecification(com.amazonaws.services.ec2.model.LaunchSpecification)
	WithLaunchSpecification(a ServicesEc2ModelLaunchSpecificationInterface) *ServicesEc2ModelRequestSpotInstancesRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.RequestSpotInstancesRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest clone()
	Clone3() *ServicesEc2ModelRequestSpotInstancesRequest
}

type ServicesEc2ModelRequestSpotInstancesRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest()
func NewServicesEc2ModelRequestSpotInstancesRequest() (*ServicesEc2ModelRequestSpotInstancesRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RequestSpotInstancesRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest(java.lang.String)
func NewServicesEc2ModelRequestSpotInstancesRequest2(a string) (*ServicesEc2ModelRequestSpotInstancesRequest) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RequestSpotInstancesRequest", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSpotPrice(java.lang.String)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) SetSpotPrice(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotPrice", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSpotPrice()
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) GetSpotPrice() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotPrice", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withSpotPrice(java.lang.String)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) WithSpotPrice(a string) *ServicesEc2ModelRequestSpotInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotPrice", "com/amazonaws/services/ec2/model/RequestSpotInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) SetClientToken(a string)  {
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
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) GetClientToken() string {
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

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) WithClientToken(a string) *ServicesEc2ModelRequestSpotInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/RequestSpotInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) SetInstanceCount(a int)  {
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
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) GetInstanceCount() int {
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

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) WithInstanceCount(a int) *ServicesEc2ModelRequestSpotInstancesRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceCount", "com/amazonaws/services/ec2/model/RequestSpotInstancesRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setType(java.lang.String)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) SetType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getType()
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) GetType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withType(java.lang.String)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) WithType2(a string) *ServicesEc2ModelRequestSpotInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withType", "com/amazonaws/services/ec2/model/RequestSpotInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setType(com.amazonaws.services.ec2.model.SpotInstanceType)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) SetType(a ServicesEc2ModelSpotInstanceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotInstanceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withType(com.amazonaws.services.ec2.model.SpotInstanceType)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) WithType(a ServicesEc2ModelSpotInstanceTypeInterface) *ServicesEc2ModelRequestSpotInstancesRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withType", "com/amazonaws/services/ec2/model/RequestSpotInstancesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotInstanceType"))
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
	unique_x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setValidFrom(java.util.Date)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) SetValidFrom(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValidFrom", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getValidFrom()
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) GetValidFrom() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getValidFrom", "java/util/Date")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDate()
	dst := new(time.Time)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withValidFrom(java.util.Date)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) WithValidFrom(a time.Time) *ServicesEc2ModelRequestSpotInstancesRequest {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withValidFrom", "com/amazonaws/services/ec2/model/RequestSpotInstancesRequest", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setValidUntil(java.util.Date)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) SetValidUntil(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValidUntil", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getValidUntil()
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) GetValidUntil() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getValidUntil", "java/util/Date")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDate()
	dst := new(time.Time)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withValidUntil(java.util.Date)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) WithValidUntil(a time.Time) *ServicesEc2ModelRequestSpotInstancesRequest {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withValidUntil", "com/amazonaws/services/ec2/model/RequestSpotInstancesRequest", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLaunchGroup(java.lang.String)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) SetLaunchGroup(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLaunchGroup", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getLaunchGroup()
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) GetLaunchGroup() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLaunchGroup", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withLaunchGroup(java.lang.String)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) WithLaunchGroup(a string) *ServicesEc2ModelRequestSpotInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLaunchGroup", "com/amazonaws/services/ec2/model/RequestSpotInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZoneGroup(java.lang.String)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) SetAvailabilityZoneGroup(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailabilityZoneGroup", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAvailabilityZoneGroup()
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) GetAvailabilityZoneGroup() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailabilityZoneGroup", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withAvailabilityZoneGroup(java.lang.String)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) WithAvailabilityZoneGroup(a string) *ServicesEc2ModelRequestSpotInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZoneGroup", "com/amazonaws/services/ec2/model/RequestSpotInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setBlockDurationMinutes(java.lang.Integer)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) SetBlockDurationMinutes(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBlockDurationMinutes", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getBlockDurationMinutes()
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) GetBlockDurationMinutes() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlockDurationMinutes", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withBlockDurationMinutes(java.lang.Integer)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) WithBlockDurationMinutes(a int) *ServicesEc2ModelRequestSpotInstancesRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDurationMinutes", "com/amazonaws/services/ec2/model/RequestSpotInstancesRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLaunchSpecification(com.amazonaws.services.ec2.model.LaunchSpecification)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) SetLaunchSpecification(a ServicesEc2ModelLaunchSpecificationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLaunchSpecification", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/LaunchSpecification"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.LaunchSpecification getLaunchSpecification()
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) GetLaunchSpecification() *ServicesEc2ModelLaunchSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLaunchSpecification", "com/amazonaws/services/ec2/model/LaunchSpecification")
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest withLaunchSpecification(com.amazonaws.services.ec2.model.LaunchSpecification)
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) WithLaunchSpecification(a ServicesEc2ModelLaunchSpecificationInterface) *ServicesEc2ModelRequestSpotInstancesRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLaunchSpecification", "com/amazonaws/services/ec2/model/RequestSpotInstancesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/LaunchSpecification"))
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
	unique_x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.RequestSpotInstancesRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.RequestSpotInstancesRequest clone()
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) Clone3() *ServicesEc2ModelRequestSpotInstancesRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/RequestSpotInstancesRequest")
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
	unique_x := &ServicesEc2ModelRequestSpotInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelRequestSpotInstancesRequest) Clone2() (*JavaLangObject, error) {
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


