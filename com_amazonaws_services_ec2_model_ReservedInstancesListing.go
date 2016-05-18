package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelReservedInstancesListingInterface interface {
	JavaLangObjectInterface

	// public void setReservedInstancesListingId(java.lang.String)
	SetReservedInstancesListingId(a string) 

	// public java.lang.String getReservedInstancesListingId()
	GetReservedInstancesListingId() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withReservedInstancesListingId(java.lang.String)
	WithReservedInstancesListingId(a string) *ServicesEc2ModelReservedInstancesListing

	// public void setReservedInstancesId(java.lang.String)
	SetReservedInstancesId(a string) 

	// public java.lang.String getReservedInstancesId()
	GetReservedInstancesId() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withReservedInstancesId(java.lang.String)
	WithReservedInstancesId(a string) *ServicesEc2ModelReservedInstancesListing

	// public void setCreateDate(java.util.Date)
	SetCreateDate(a time.Time) 

	// public java.util.Date getCreateDate()
	GetCreateDate() time.Time

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withCreateDate(java.util.Date)
	WithCreateDate(a time.Time) *ServicesEc2ModelReservedInstancesListing

	// public void setUpdateDate(java.util.Date)
	SetUpdateDate(a time.Time) 

	// public java.util.Date getUpdateDate()
	GetUpdateDate() time.Time

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withUpdateDate(java.util.Date)
	WithUpdateDate(a time.Time) *ServicesEc2ModelReservedInstancesListing

	// public void setStatus(java.lang.String)
	SetStatus2(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withStatus(java.lang.String)
	WithStatus2(a string) *ServicesEc2ModelReservedInstancesListing

	// public void setStatus(com.amazonaws.services.ec2.model.ListingStatus)
	SetStatus(a ServicesEc2ModelListingStatusInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withStatus(com.amazonaws.services.ec2.model.ListingStatus)
	WithStatus(a ServicesEc2ModelListingStatusInterface) *ServicesEc2ModelReservedInstancesListing

	// public void setStatusMessage(java.lang.String)
	SetStatusMessage(a string) 

	// public java.lang.String getStatusMessage()
	GetStatusMessage() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withStatusMessage(java.lang.String)
	WithStatusMessage(a string) *ServicesEc2ModelReservedInstancesListing

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceCount> getInstanceCounts()
	GetInstanceCounts() []*ServicesEc2ModelInstanceCount

	// public void setInstanceCounts(java.util.Collection<com.amazonaws.services.ec2.model.InstanceCount>)
	SetInstanceCounts(a []*ServicesEc2ModelInstanceCount) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withInstanceCounts(com.amazonaws.services.ec2.model.InstanceCount...)
	WithInstanceCounts(a ...*ServicesEc2ModelInstanceCount) *ServicesEc2ModelReservedInstancesListing

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withInstanceCounts(java.util.Collection<com.amazonaws.services.ec2.model.InstanceCount>)
	WithInstanceCounts2(a []*ServicesEc2ModelInstanceCount) *ServicesEc2ModelReservedInstancesListing

	// public java.util.List<com.amazonaws.services.ec2.model.PriceSchedule> getPriceSchedules()
	GetPriceSchedules() []*ServicesEc2ModelPriceSchedule

	// public void setPriceSchedules(java.util.Collection<com.amazonaws.services.ec2.model.PriceSchedule>)
	SetPriceSchedules(a []*ServicesEc2ModelPriceSchedule) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withPriceSchedules(com.amazonaws.services.ec2.model.PriceSchedule...)
	WithPriceSchedules(a ...*ServicesEc2ModelPriceSchedule) *ServicesEc2ModelReservedInstancesListing

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withPriceSchedules(java.util.Collection<com.amazonaws.services.ec2.model.PriceSchedule>)
	WithPriceSchedules2(a []*ServicesEc2ModelPriceSchedule) *ServicesEc2ModelReservedInstancesListing

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelReservedInstancesListing

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelReservedInstancesListing

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelReservedInstancesListing

	// public com.amazonaws.services.ec2.model.ReservedInstancesListing clone()
	Clone() *ServicesEc2ModelReservedInstancesListing
}

type ServicesEc2ModelReservedInstancesListing struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ReservedInstancesListing()
func NewServicesEc2ModelReservedInstancesListing() (*ServicesEc2ModelReservedInstancesListing) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ReservedInstancesListing")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelReservedInstancesListing{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setReservedInstancesListingId(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesListing) SetReservedInstancesListingId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesListingId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getReservedInstancesListingId()
func (jbobject *ServicesEc2ModelReservedInstancesListing) GetReservedInstancesListingId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesListingId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withReservedInstancesListingId(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithReservedInstancesListingId(a string) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesListingId", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public void setReservedInstancesId(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesListing) SetReservedInstancesId(a string)  {
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
func (jbobject *ServicesEc2ModelReservedInstancesListing) GetReservedInstancesId() string {
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

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withReservedInstancesId(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithReservedInstancesId(a string) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesId", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCreateDate(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstancesListing) SetCreateDate(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCreateDate", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getCreateDate()
func (jbobject *ServicesEc2ModelReservedInstancesListing) GetCreateDate() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCreateDate", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withCreateDate(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithCreateDate(a time.Time) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCreateDate", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUpdateDate(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstancesListing) SetUpdateDate(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUpdateDate", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getUpdateDate()
func (jbobject *ServicesEc2ModelReservedInstancesListing) GetUpdateDate() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUpdateDate", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withUpdateDate(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithUpdateDate(a time.Time) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUpdateDate", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesListing) SetStatus2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatus()
func (jbobject *ServicesEc2ModelReservedInstancesListing) GetStatus() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatus", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithStatus2(a string) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(com.amazonaws.services.ec2.model.ListingStatus)
func (jbobject *ServicesEc2ModelReservedInstancesListing) SetStatus(a ServicesEc2ModelListingStatusInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ListingStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withStatus(com.amazonaws.services.ec2.model.ListingStatus)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithStatus(a ServicesEc2ModelListingStatusInterface) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ListingStatus"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesListing) SetStatusMessage(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatusMessage", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatusMessage()
func (jbobject *ServicesEc2ModelReservedInstancesListing) GetStatusMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatusMessage", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithStatusMessage(a string) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatusMessage", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceCount> getInstanceCounts()
func (jbobject *ServicesEc2ModelReservedInstancesListing) GetInstanceCounts() []*ServicesEc2ModelInstanceCount {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceCounts", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInstanceCount)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setInstanceCounts(java.util.Collection<com.amazonaws.services.ec2.model.InstanceCount>)
func (jbobject *ServicesEc2ModelReservedInstancesListing) SetInstanceCounts(a []*ServicesEc2ModelInstanceCount)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceCounts", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withInstanceCounts(com.amazonaws.services.ec2.model.InstanceCount...)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithInstanceCounts(a ...*ServicesEc2ModelInstanceCount) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceCount")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceCounts", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceCount"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withInstanceCounts(java.util.Collection<com.amazonaws.services.ec2.model.InstanceCount>)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithInstanceCounts2(a []*ServicesEc2ModelInstanceCount) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceCounts", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.PriceSchedule> getPriceSchedules()
func (jbobject *ServicesEc2ModelReservedInstancesListing) GetPriceSchedules() []*ServicesEc2ModelPriceSchedule {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPriceSchedules", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelPriceSchedule)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setPriceSchedules(java.util.Collection<com.amazonaws.services.ec2.model.PriceSchedule>)
func (jbobject *ServicesEc2ModelReservedInstancesListing) SetPriceSchedules(a []*ServicesEc2ModelPriceSchedule)  {
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

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withPriceSchedules(com.amazonaws.services.ec2.model.PriceSchedule...)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithPriceSchedules(a ...*ServicesEc2ModelPriceSchedule) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/PriceSchedule")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPriceSchedules", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PriceSchedule"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withPriceSchedules(java.util.Collection<com.amazonaws.services.ec2.model.PriceSchedule>)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithPriceSchedules2(a []*ServicesEc2ModelPriceSchedule) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPriceSchedules", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelReservedInstancesListing) GetTags() []*ServicesEc2ModelTag {
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
func (jbobject *ServicesEc2ModelReservedInstancesListing) SetTags(a []*ServicesEc2ModelTag)  {
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

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesListing) SetClientToken(a string)  {
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
func (jbobject *ServicesEc2ModelReservedInstancesListing) GetClientToken() string {
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

// public com.amazonaws.services.ec2.model.ReservedInstancesListing withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesListing) WithClientToken(a string) *ServicesEc2ModelReservedInstancesListing {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/ReservedInstancesListing", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelReservedInstancesListing) ToString() string {
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
func (jbobject *ServicesEc2ModelReservedInstancesListing) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelReservedInstancesListing) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ReservedInstancesListing clone()
func (jbobject *ServicesEc2ModelReservedInstancesListing) Clone() *ServicesEc2ModelReservedInstancesListing {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ReservedInstancesListing")
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
	unique_x := &ServicesEc2ModelReservedInstancesListing{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelReservedInstancesListing) Clone2() (*JavaLangObject, error) {
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


