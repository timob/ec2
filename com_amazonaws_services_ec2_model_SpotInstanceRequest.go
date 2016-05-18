package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelSpotInstanceRequestInterface interface {
	JavaLangObjectInterface

	// public void setSpotInstanceRequestId(java.lang.String)
	SetSpotInstanceRequestId(a string) 

	// public java.lang.String getSpotInstanceRequestId()
	GetSpotInstanceRequestId() string

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withSpotInstanceRequestId(java.lang.String)
	WithSpotInstanceRequestId(a string) *ServicesEc2ModelSpotInstanceRequest

	// public void setSpotPrice(java.lang.String)
	SetSpotPrice(a string) 

	// public java.lang.String getSpotPrice()
	GetSpotPrice() string

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withSpotPrice(java.lang.String)
	WithSpotPrice(a string) *ServicesEc2ModelSpotInstanceRequest

	// public void setType(java.lang.String)
	SetType2(a string) 

	// public java.lang.String getType()
	GetType() string

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withType(java.lang.String)
	WithType2(a string) *ServicesEc2ModelSpotInstanceRequest

	// public void setType(com.amazonaws.services.ec2.model.SpotInstanceType)
	SetType(a ServicesEc2ModelSpotInstanceTypeInterface) 

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withType(com.amazonaws.services.ec2.model.SpotInstanceType)
	WithType(a ServicesEc2ModelSpotInstanceTypeInterface) *ServicesEc2ModelSpotInstanceRequest

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelSpotInstanceRequest

	// public void setState(com.amazonaws.services.ec2.model.SpotInstanceState)
	SetState(a ServicesEc2ModelSpotInstanceStateInterface) 

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withState(com.amazonaws.services.ec2.model.SpotInstanceState)
	WithState(a ServicesEc2ModelSpotInstanceStateInterface) *ServicesEc2ModelSpotInstanceRequest

	// public void setFault(com.amazonaws.services.ec2.model.SpotInstanceStateFault)
	SetFault(a ServicesEc2ModelSpotInstanceStateFaultInterface) 

	// public com.amazonaws.services.ec2.model.SpotInstanceStateFault getFault()
	GetFault() *ServicesEc2ModelSpotInstanceStateFault

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withFault(com.amazonaws.services.ec2.model.SpotInstanceStateFault)
	WithFault(a ServicesEc2ModelSpotInstanceStateFaultInterface) *ServicesEc2ModelSpotInstanceRequest

	// public void setStatus(com.amazonaws.services.ec2.model.SpotInstanceStatus)
	SetStatus(a ServicesEc2ModelSpotInstanceStatusInterface) 

	// public com.amazonaws.services.ec2.model.SpotInstanceStatus getStatus()
	GetStatus() *ServicesEc2ModelSpotInstanceStatus

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withStatus(com.amazonaws.services.ec2.model.SpotInstanceStatus)
	WithStatus(a ServicesEc2ModelSpotInstanceStatusInterface) *ServicesEc2ModelSpotInstanceRequest

	// public void setValidFrom(java.util.Date)
	SetValidFrom(a time.Time) 

	// public java.util.Date getValidFrom()
	GetValidFrom() time.Time

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withValidFrom(java.util.Date)
	WithValidFrom(a time.Time) *ServicesEc2ModelSpotInstanceRequest

	// public void setValidUntil(java.util.Date)
	SetValidUntil(a time.Time) 

	// public java.util.Date getValidUntil()
	GetValidUntil() time.Time

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withValidUntil(java.util.Date)
	WithValidUntil(a time.Time) *ServicesEc2ModelSpotInstanceRequest

	// public void setLaunchGroup(java.lang.String)
	SetLaunchGroup(a string) 

	// public java.lang.String getLaunchGroup()
	GetLaunchGroup() string

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withLaunchGroup(java.lang.String)
	WithLaunchGroup(a string) *ServicesEc2ModelSpotInstanceRequest

	// public void setAvailabilityZoneGroup(java.lang.String)
	SetAvailabilityZoneGroup(a string) 

	// public java.lang.String getAvailabilityZoneGroup()
	GetAvailabilityZoneGroup() string

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withAvailabilityZoneGroup(java.lang.String)
	WithAvailabilityZoneGroup(a string) *ServicesEc2ModelSpotInstanceRequest

	// public void setLaunchSpecification(com.amazonaws.services.ec2.model.LaunchSpecification)
	SetLaunchSpecification(a ServicesEc2ModelLaunchSpecificationInterface) 

	// public com.amazonaws.services.ec2.model.LaunchSpecification getLaunchSpecification()
	GetLaunchSpecification() *ServicesEc2ModelLaunchSpecification

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withLaunchSpecification(com.amazonaws.services.ec2.model.LaunchSpecification)
	WithLaunchSpecification(a ServicesEc2ModelLaunchSpecificationInterface) *ServicesEc2ModelSpotInstanceRequest

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelSpotInstanceRequest

	// public void setCreateTime(java.util.Date)
	SetCreateTime(a time.Time) 

	// public java.util.Date getCreateTime()
	GetCreateTime() time.Time

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withCreateTime(java.util.Date)
	WithCreateTime(a time.Time) *ServicesEc2ModelSpotInstanceRequest

	// public void setProductDescription(java.lang.String)
	SetProductDescription2(a string) 

	// public java.lang.String getProductDescription()
	GetProductDescription() string

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withProductDescription(java.lang.String)
	WithProductDescription2(a string) *ServicesEc2ModelSpotInstanceRequest

	// public void setProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
	SetProductDescription(a ServicesEc2ModelRIProductDescriptionInterface) 

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
	WithProductDescription(a ServicesEc2ModelRIProductDescriptionInterface) *ServicesEc2ModelSpotInstanceRequest

	// public void setBlockDurationMinutes(java.lang.Integer)
	SetBlockDurationMinutes(a int) 

	// public java.lang.Integer getBlockDurationMinutes()
	GetBlockDurationMinutes() int

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withBlockDurationMinutes(java.lang.Integer)
	WithBlockDurationMinutes(a int) *ServicesEc2ModelSpotInstanceRequest

	// public void setActualBlockHourlyPrice(java.lang.String)
	SetActualBlockHourlyPrice(a string) 

	// public java.lang.String getActualBlockHourlyPrice()
	GetActualBlockHourlyPrice() string

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withActualBlockHourlyPrice(java.lang.String)
	WithActualBlockHourlyPrice(a string) *ServicesEc2ModelSpotInstanceRequest

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelSpotInstanceRequest

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelSpotInstanceRequest

	// public void setLaunchedAvailabilityZone(java.lang.String)
	SetLaunchedAvailabilityZone(a string) 

	// public java.lang.String getLaunchedAvailabilityZone()
	GetLaunchedAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest withLaunchedAvailabilityZone(java.lang.String)
	WithLaunchedAvailabilityZone(a string) *ServicesEc2ModelSpotInstanceRequest

	// public com.amazonaws.services.ec2.model.SpotInstanceRequest clone()
	Clone() *ServicesEc2ModelSpotInstanceRequest
}

type ServicesEc2ModelSpotInstanceRequest struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.SpotInstanceRequest()
func NewServicesEc2ModelSpotInstanceRequest() (*ServicesEc2ModelSpotInstanceRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/SpotInstanceRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelSpotInstanceRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSpotInstanceRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetSpotInstanceRequestId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotInstanceRequestId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSpotInstanceRequestId()
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetSpotInstanceRequestId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotInstanceRequestId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withSpotInstanceRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithSpotInstanceRequestId(a string) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotInstanceRequestId", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSpotPrice(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetSpotPrice(a string)  {
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
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetSpotPrice() string {
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withSpotPrice(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithSpotPrice(a string) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotPrice", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setType(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetType2(a string)  {
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
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetType() string {
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withType(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithType2(a string) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withType", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setType(com.amazonaws.services.ec2.model.SpotInstanceType)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetType(a ServicesEc2ModelSpotInstanceTypeInterface)  {
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withType(com.amazonaws.services.ec2.model.SpotInstanceType)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithType(a ServicesEc2ModelSpotInstanceTypeInterface) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withType", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotInstanceType"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetState2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getState()
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetState() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getState", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withState(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithState2(a string) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.SpotInstanceState)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetState(a ServicesEc2ModelSpotInstanceStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotInstanceState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withState(com.amazonaws.services.ec2.model.SpotInstanceState)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithState(a ServicesEc2ModelSpotInstanceStateInterface) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotInstanceState"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFault(com.amazonaws.services.ec2.model.SpotInstanceStateFault)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetFault(a ServicesEc2ModelSpotInstanceStateFaultInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFault", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotInstanceStateFault"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotInstanceStateFault getFault()
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetFault() *ServicesEc2ModelSpotInstanceStateFault {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFault", "com/amazonaws/services/ec2/model/SpotInstanceStateFault")
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
	unique_x := &ServicesEc2ModelSpotInstanceStateFault{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withFault(com.amazonaws.services.ec2.model.SpotInstanceStateFault)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithFault(a ServicesEc2ModelSpotInstanceStateFaultInterface) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFault", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotInstanceStateFault"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(com.amazonaws.services.ec2.model.SpotInstanceStatus)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetStatus(a ServicesEc2ModelSpotInstanceStatusInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotInstanceStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotInstanceStatus getStatus()
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetStatus() *ServicesEc2ModelSpotInstanceStatus {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatus", "com/amazonaws/services/ec2/model/SpotInstanceStatus")
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
	unique_x := &ServicesEc2ModelSpotInstanceStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withStatus(com.amazonaws.services.ec2.model.SpotInstanceStatus)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithStatus(a ServicesEc2ModelSpotInstanceStatusInterface) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotInstanceStatus"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setValidFrom(java.util.Date)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetValidFrom(a time.Time)  {
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
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetValidFrom() time.Time {
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withValidFrom(java.util.Date)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithValidFrom(a time.Time) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withValidFrom", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setValidUntil(java.util.Date)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetValidUntil(a time.Time)  {
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
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetValidUntil() time.Time {
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withValidUntil(java.util.Date)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithValidUntil(a time.Time) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withValidUntil", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLaunchGroup(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetLaunchGroup(a string)  {
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
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetLaunchGroup() string {
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withLaunchGroup(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithLaunchGroup(a string) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLaunchGroup", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZoneGroup(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetAvailabilityZoneGroup(a string)  {
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
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetAvailabilityZoneGroup() string {
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withAvailabilityZoneGroup(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithAvailabilityZoneGroup(a string) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZoneGroup", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLaunchSpecification(com.amazonaws.services.ec2.model.LaunchSpecification)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetLaunchSpecification(a ServicesEc2ModelLaunchSpecificationInterface)  {
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
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetLaunchSpecification() *ServicesEc2ModelLaunchSpecification {
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withLaunchSpecification(com.amazonaws.services.ec2.model.LaunchSpecification)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithLaunchSpecification(a ServicesEc2ModelLaunchSpecificationInterface) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLaunchSpecification", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/LaunchSpecification"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceId()
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithInstanceId(a string) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCreateTime(java.util.Date)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetCreateTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCreateTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getCreateTime()
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetCreateTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCreateTime", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withCreateTime(java.util.Date)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithCreateTime(a time.Time) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCreateTime", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProductDescription(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetProductDescription2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductDescription", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getProductDescription()
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetProductDescription() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProductDescription", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withProductDescription(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithProductDescription2(a string) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductDescription", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetProductDescription(a ServicesEc2ModelRIProductDescriptionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductDescription", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/RIProductDescription"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithProductDescription(a ServicesEc2ModelRIProductDescriptionInterface) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductDescription", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RIProductDescription"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setBlockDurationMinutes(java.lang.Integer)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetBlockDurationMinutes(a int)  {
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
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetBlockDurationMinutes() int {
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withBlockDurationMinutes(java.lang.Integer)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithBlockDurationMinutes(a int) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDurationMinutes", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setActualBlockHourlyPrice(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetActualBlockHourlyPrice(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setActualBlockHourlyPrice", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getActualBlockHourlyPrice()
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetActualBlockHourlyPrice() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getActualBlockHourlyPrice", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withActualBlockHourlyPrice(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithActualBlockHourlyPrice(a string) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withActualBlockHourlyPrice", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetTags() []*ServicesEc2ModelTag {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTags", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelTag)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetTags(a []*ServicesEc2ModelTag)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTags", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLaunchedAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) SetLaunchedAvailabilityZone(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLaunchedAvailabilityZone", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getLaunchedAvailabilityZone()
func (jbobject *ServicesEc2ModelSpotInstanceRequest) GetLaunchedAvailabilityZone() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLaunchedAvailabilityZone", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotInstanceRequest withLaunchedAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceRequest) WithLaunchedAvailabilityZone(a string) *ServicesEc2ModelSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLaunchedAvailabilityZone", "com/amazonaws/services/ec2/model/SpotInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSpotInstanceRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelSpotInstanceRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelSpotInstanceRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.SpotInstanceRequest clone()
func (jbobject *ServicesEc2ModelSpotInstanceRequest) Clone() *ServicesEc2ModelSpotInstanceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/SpotInstanceRequest")
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
	unique_x := &ServicesEc2ModelSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelSpotInstanceRequest) Clone2() (*JavaLangObject, error) {
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


