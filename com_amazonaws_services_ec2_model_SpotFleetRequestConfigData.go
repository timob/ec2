package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelSpotFleetRequestConfigDataInterface interface {
	JavaLangObjectInterface

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelSpotFleetRequestConfigData

	// public void setSpotPrice(java.lang.String)
	SetSpotPrice(a string) 

	// public java.lang.String getSpotPrice()
	GetSpotPrice() string

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withSpotPrice(java.lang.String)
	WithSpotPrice(a string) *ServicesEc2ModelSpotFleetRequestConfigData

	// public void setTargetCapacity(java.lang.Integer)
	SetTargetCapacity(a int) 

	// public java.lang.Integer getTargetCapacity()
	GetTargetCapacity() int

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withTargetCapacity(java.lang.Integer)
	WithTargetCapacity(a int) *ServicesEc2ModelSpotFleetRequestConfigData

	// public void setValidFrom(java.util.Date)
	SetValidFrom(a time.Time) 

	// public java.util.Date getValidFrom()
	GetValidFrom() time.Time

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withValidFrom(java.util.Date)
	WithValidFrom(a time.Time) *ServicesEc2ModelSpotFleetRequestConfigData

	// public void setValidUntil(java.util.Date)
	SetValidUntil(a time.Time) 

	// public java.util.Date getValidUntil()
	GetValidUntil() time.Time

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withValidUntil(java.util.Date)
	WithValidUntil(a time.Time) *ServicesEc2ModelSpotFleetRequestConfigData

	// public void setTerminateInstancesWithExpiration(java.lang.Boolean)
	SetTerminateInstancesWithExpiration(a bool) 

	// public java.lang.Boolean getTerminateInstancesWithExpiration()
	GetTerminateInstancesWithExpiration() bool

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withTerminateInstancesWithExpiration(java.lang.Boolean)
	WithTerminateInstancesWithExpiration(a bool) *ServicesEc2ModelSpotFleetRequestConfigData

	// public java.lang.Boolean isTerminateInstancesWithExpiration()
	IsTerminateInstancesWithExpiration() bool

	// public void setIamFleetRole(java.lang.String)
	SetIamFleetRole(a string) 

	// public java.lang.String getIamFleetRole()
	GetIamFleetRole() string

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withIamFleetRole(java.lang.String)
	WithIamFleetRole(a string) *ServicesEc2ModelSpotFleetRequestConfigData

	// public java.util.List<com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification> getLaunchSpecifications()
	GetLaunchSpecifications() []*ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setLaunchSpecifications(java.util.Collection<com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification>)
	SetLaunchSpecifications(a []*ServicesEc2ModelSpotFleetLaunchSpecification) 

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withLaunchSpecifications(com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification...)
	WithLaunchSpecifications(a ...*ServicesEc2ModelSpotFleetLaunchSpecification) *ServicesEc2ModelSpotFleetRequestConfigData

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withLaunchSpecifications(java.util.Collection<com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification>)
	WithLaunchSpecifications2(a []*ServicesEc2ModelSpotFleetLaunchSpecification) *ServicesEc2ModelSpotFleetRequestConfigData

	// public void setExcessCapacityTerminationPolicy(java.lang.String)
	SetExcessCapacityTerminationPolicy2(a string) 

	// public java.lang.String getExcessCapacityTerminationPolicy()
	GetExcessCapacityTerminationPolicy() string

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withExcessCapacityTerminationPolicy(java.lang.String)
	WithExcessCapacityTerminationPolicy2(a string) *ServicesEc2ModelSpotFleetRequestConfigData

	// public void setExcessCapacityTerminationPolicy(com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy)
	SetExcessCapacityTerminationPolicy(a ServicesEc2ModelExcessCapacityTerminationPolicyInterface) 

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withExcessCapacityTerminationPolicy(com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy)
	WithExcessCapacityTerminationPolicy(a ServicesEc2ModelExcessCapacityTerminationPolicyInterface) *ServicesEc2ModelSpotFleetRequestConfigData

	// public void setAllocationStrategy(java.lang.String)
	SetAllocationStrategy2(a string) 

	// public java.lang.String getAllocationStrategy()
	GetAllocationStrategy() string

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withAllocationStrategy(java.lang.String)
	WithAllocationStrategy2(a string) *ServicesEc2ModelSpotFleetRequestConfigData

	// public void setAllocationStrategy(com.amazonaws.services.ec2.model.AllocationStrategy)
	SetAllocationStrategy(a ServicesEc2ModelAllocationStrategyInterface) 

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withAllocationStrategy(com.amazonaws.services.ec2.model.AllocationStrategy)
	WithAllocationStrategy(a ServicesEc2ModelAllocationStrategyInterface) *ServicesEc2ModelSpotFleetRequestConfigData

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData clone()
	Clone() *ServicesEc2ModelSpotFleetRequestConfigData
}

type ServicesEc2ModelSpotFleetRequestConfigData struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData()
func NewServicesEc2ModelSpotFleetRequestConfigData() (*ServicesEc2ModelSpotFleetRequestConfigData) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/SpotFleetRequestConfigData")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) SetClientToken(a string)  {
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
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) GetClientToken() string {
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

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) WithClientToken(a string) *ServicesEc2ModelSpotFleetRequestConfigData {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSpotPrice(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) SetSpotPrice(a string)  {
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
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) GetSpotPrice() string {
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

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withSpotPrice(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) WithSpotPrice(a string) *ServicesEc2ModelSpotFleetRequestConfigData {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotPrice", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTargetCapacity(java.lang.Integer)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) SetTargetCapacity(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTargetCapacity", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getTargetCapacity()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) GetTargetCapacity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTargetCapacity", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withTargetCapacity(java.lang.Integer)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) WithTargetCapacity(a int) *ServicesEc2ModelSpotFleetRequestConfigData {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTargetCapacity", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public void setValidFrom(java.util.Date)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) SetValidFrom(a time.Time)  {
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
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) GetValidFrom() time.Time {
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

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withValidFrom(java.util.Date)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) WithValidFrom(a time.Time) *ServicesEc2ModelSpotFleetRequestConfigData {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withValidFrom", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public void setValidUntil(java.util.Date)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) SetValidUntil(a time.Time)  {
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
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) GetValidUntil() time.Time {
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

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withValidUntil(java.util.Date)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) WithValidUntil(a time.Time) *ServicesEc2ModelSpotFleetRequestConfigData {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withValidUntil", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTerminateInstancesWithExpiration(java.lang.Boolean)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) SetTerminateInstancesWithExpiration(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTerminateInstancesWithExpiration", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getTerminateInstancesWithExpiration()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) GetTerminateInstancesWithExpiration() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTerminateInstancesWithExpiration", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withTerminateInstancesWithExpiration(java.lang.Boolean)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) WithTerminateInstancesWithExpiration(a bool) *ServicesEc2ModelSpotFleetRequestConfigData {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTerminateInstancesWithExpiration", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isTerminateInstancesWithExpiration()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) IsTerminateInstancesWithExpiration() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTerminateInstancesWithExpiration", "java/lang/Boolean")
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

// public void setIamFleetRole(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) SetIamFleetRole(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIamFleetRole", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getIamFleetRole()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) GetIamFleetRole() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIamFleetRole", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withIamFleetRole(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) WithIamFleetRole(a string) *ServicesEc2ModelSpotFleetRequestConfigData {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIamFleetRole", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification> getLaunchSpecifications()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) GetLaunchSpecifications() []*ServicesEc2ModelSpotFleetLaunchSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLaunchSpecifications", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelSpotFleetLaunchSpecification)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setLaunchSpecifications(java.util.Collection<com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification>)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) SetLaunchSpecifications(a []*ServicesEc2ModelSpotFleetLaunchSpecification)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLaunchSpecifications", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withLaunchSpecifications(com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification...)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) WithLaunchSpecifications(a ...*ServicesEc2ModelSpotFleetLaunchSpecification) *ServicesEc2ModelSpotFleetRequestConfigData {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLaunchSpecifications", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withLaunchSpecifications(java.util.Collection<com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification>)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) WithLaunchSpecifications2(a []*ServicesEc2ModelSpotFleetLaunchSpecification) *ServicesEc2ModelSpotFleetRequestConfigData {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLaunchSpecifications", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public void setExcessCapacityTerminationPolicy(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) SetExcessCapacityTerminationPolicy2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExcessCapacityTerminationPolicy", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getExcessCapacityTerminationPolicy()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) GetExcessCapacityTerminationPolicy() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExcessCapacityTerminationPolicy", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withExcessCapacityTerminationPolicy(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) WithExcessCapacityTerminationPolicy2(a string) *ServicesEc2ModelSpotFleetRequestConfigData {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withExcessCapacityTerminationPolicy", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public void setExcessCapacityTerminationPolicy(com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) SetExcessCapacityTerminationPolicy(a ServicesEc2ModelExcessCapacityTerminationPolicyInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExcessCapacityTerminationPolicy", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withExcessCapacityTerminationPolicy(com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) WithExcessCapacityTerminationPolicy(a ServicesEc2ModelExcessCapacityTerminationPolicyInterface) *ServicesEc2ModelSpotFleetRequestConfigData {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withExcessCapacityTerminationPolicy", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAllocationStrategy(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) SetAllocationStrategy2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAllocationStrategy", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAllocationStrategy()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) GetAllocationStrategy() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAllocationStrategy", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withAllocationStrategy(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) WithAllocationStrategy2(a string) *ServicesEc2ModelSpotFleetRequestConfigData {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAllocationStrategy", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAllocationStrategy(com.amazonaws.services.ec2.model.AllocationStrategy)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) SetAllocationStrategy(a ServicesEc2ModelAllocationStrategyInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAllocationStrategy", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AllocationStrategy"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData withAllocationStrategy(com.amazonaws.services.ec2.model.AllocationStrategy)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) WithAllocationStrategy(a ServicesEc2ModelAllocationStrategyInterface) *ServicesEc2ModelSpotFleetRequestConfigData {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAllocationStrategy", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AllocationStrategy"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) ToString() string {
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
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData clone()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) Clone() *ServicesEc2ModelSpotFleetRequestConfigData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData")
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelSpotFleetRequestConfigData) Clone2() (*JavaLangObject, error) {
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


