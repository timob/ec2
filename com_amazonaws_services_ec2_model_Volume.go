package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelVolumeInterface interface {
	JavaLangObjectInterface

	// public void setVolumeId(java.lang.String)
	SetVolumeId(a string) 

	// public java.lang.String getVolumeId()
	GetVolumeId() string

	// public com.amazonaws.services.ec2.model.Volume withVolumeId(java.lang.String)
	WithVolumeId(a string) *ServicesEc2ModelVolume

	// public void setSize(java.lang.Integer)
	SetSize(a int) 

	// public java.lang.Integer getSize()
	GetSize() int

	// public com.amazonaws.services.ec2.model.Volume withSize(java.lang.Integer)
	WithSize(a int) *ServicesEc2ModelVolume

	// public void setSnapshotId(java.lang.String)
	SetSnapshotId(a string) 

	// public java.lang.String getSnapshotId()
	GetSnapshotId() string

	// public com.amazonaws.services.ec2.model.Volume withSnapshotId(java.lang.String)
	WithSnapshotId(a string) *ServicesEc2ModelVolume

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.Volume withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelVolume

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.Volume withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelVolume

	// public void setState(com.amazonaws.services.ec2.model.VolumeState)
	SetState(a ServicesEc2ModelVolumeStateInterface) 

	// public com.amazonaws.services.ec2.model.Volume withState(com.amazonaws.services.ec2.model.VolumeState)
	WithState(a ServicesEc2ModelVolumeStateInterface) *ServicesEc2ModelVolume

	// public void setCreateTime(java.util.Date)
	SetCreateTime(a time.Time) 

	// public java.util.Date getCreateTime()
	GetCreateTime() time.Time

	// public com.amazonaws.services.ec2.model.Volume withCreateTime(java.util.Date)
	WithCreateTime(a time.Time) *ServicesEc2ModelVolume

	// public java.util.List<com.amazonaws.services.ec2.model.VolumeAttachment> getAttachments()
	GetAttachments() []*ServicesEc2ModelVolumeAttachment

	// public void setAttachments(java.util.Collection<com.amazonaws.services.ec2.model.VolumeAttachment>)
	SetAttachments(a []*ServicesEc2ModelVolumeAttachment) 

	// public com.amazonaws.services.ec2.model.Volume withAttachments(com.amazonaws.services.ec2.model.VolumeAttachment...)
	WithAttachments(a ...*ServicesEc2ModelVolumeAttachment) *ServicesEc2ModelVolume

	// public com.amazonaws.services.ec2.model.Volume withAttachments(java.util.Collection<com.amazonaws.services.ec2.model.VolumeAttachment>)
	WithAttachments2(a []*ServicesEc2ModelVolumeAttachment) *ServicesEc2ModelVolume

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.Volume withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelVolume

	// public com.amazonaws.services.ec2.model.Volume withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelVolume

	// public void setVolumeType(java.lang.String)
	SetVolumeType2(a string) 

	// public java.lang.String getVolumeType()
	GetVolumeType() string

	// public com.amazonaws.services.ec2.model.Volume withVolumeType(java.lang.String)
	WithVolumeType2(a string) *ServicesEc2ModelVolume

	// public void setVolumeType(com.amazonaws.services.ec2.model.VolumeType)
	SetVolumeType(a ServicesEc2ModelVolumeTypeInterface) 

	// public com.amazonaws.services.ec2.model.Volume withVolumeType(com.amazonaws.services.ec2.model.VolumeType)
	WithVolumeType(a ServicesEc2ModelVolumeTypeInterface) *ServicesEc2ModelVolume

	// public void setIops(java.lang.Integer)
	SetIops(a int) 

	// public java.lang.Integer getIops()
	GetIops() int

	// public com.amazonaws.services.ec2.model.Volume withIops(java.lang.Integer)
	WithIops(a int) *ServicesEc2ModelVolume

	// public void setEncrypted(java.lang.Boolean)
	SetEncrypted(a bool) 

	// public java.lang.Boolean getEncrypted()
	GetEncrypted() bool

	// public com.amazonaws.services.ec2.model.Volume withEncrypted(java.lang.Boolean)
	WithEncrypted(a bool) *ServicesEc2ModelVolume

	// public java.lang.Boolean isEncrypted()
	IsEncrypted() bool

	// public void setKmsKeyId(java.lang.String)
	SetKmsKeyId(a string) 

	// public java.lang.String getKmsKeyId()
	GetKmsKeyId() string

	// public com.amazonaws.services.ec2.model.Volume withKmsKeyId(java.lang.String)
	WithKmsKeyId(a string) *ServicesEc2ModelVolume

	// public com.amazonaws.services.ec2.model.Volume clone()
	Clone() *ServicesEc2ModelVolume
}

type ServicesEc2ModelVolume struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.Volume()
func NewServicesEc2ModelVolume() (*ServicesEc2ModelVolume) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/Volume")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVolume{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVolumeId(java.lang.String)
func (jbobject *ServicesEc2ModelVolume) SetVolumeId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVolumeId()
func (jbobject *ServicesEc2ModelVolume) GetVolumeId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolumeId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Volume withVolumeId(java.lang.String)
func (jbobject *ServicesEc2ModelVolume) WithVolumeId(a string) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeId", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSize(java.lang.Integer)
func (jbobject *ServicesEc2ModelVolume) SetSize(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSize", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getSize()
func (jbobject *ServicesEc2ModelVolume) GetSize() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSize", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.Volume withSize(java.lang.Integer)
func (jbobject *ServicesEc2ModelVolume) WithSize(a int) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSize", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelVolume) SetSnapshotId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSnapshotId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSnapshotId()
func (jbobject *ServicesEc2ModelVolume) GetSnapshotId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSnapshotId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Volume withSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelVolume) WithSnapshotId(a string) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotId", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelVolume) SetAvailabilityZone(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailabilityZone", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAvailabilityZone()
func (jbobject *ServicesEc2ModelVolume) GetAvailabilityZone() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailabilityZone", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Volume withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelVolume) WithAvailabilityZone(a string) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelVolume) SetState2(a string)  {
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
func (jbobject *ServicesEc2ModelVolume) GetState() string {
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

// public com.amazonaws.services.ec2.model.Volume withState(java.lang.String)
func (jbobject *ServicesEc2ModelVolume) WithState2(a string) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.VolumeState)
func (jbobject *ServicesEc2ModelVolume) SetState(a ServicesEc2ModelVolumeStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Volume withState(com.amazonaws.services.ec2.model.VolumeState)
func (jbobject *ServicesEc2ModelVolume) WithState(a ServicesEc2ModelVolumeStateInterface) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeState"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCreateTime(java.util.Date)
func (jbobject *ServicesEc2ModelVolume) SetCreateTime(a time.Time)  {
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
func (jbobject *ServicesEc2ModelVolume) GetCreateTime() time.Time {
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

// public com.amazonaws.services.ec2.model.Volume withCreateTime(java.util.Date)
func (jbobject *ServicesEc2ModelVolume) WithCreateTime(a time.Time) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCreateTime", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.VolumeAttachment> getAttachments()
func (jbobject *ServicesEc2ModelVolume) GetAttachments() []*ServicesEc2ModelVolumeAttachment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttachments", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelVolumeAttachment)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAttachments(java.util.Collection<com.amazonaws.services.ec2.model.VolumeAttachment>)
func (jbobject *ServicesEc2ModelVolume) SetAttachments(a []*ServicesEc2ModelVolumeAttachment)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttachments", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Volume withAttachments(com.amazonaws.services.ec2.model.VolumeAttachment...)
func (jbobject *ServicesEc2ModelVolume) WithAttachments(a ...*ServicesEc2ModelVolumeAttachment) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/VolumeAttachment")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachments", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeAttachment"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Volume withAttachments(java.util.Collection<com.amazonaws.services.ec2.model.VolumeAttachment>)
func (jbobject *ServicesEc2ModelVolume) WithAttachments2(a []*ServicesEc2ModelVolumeAttachment) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachments", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelVolume) GetTags() []*ServicesEc2ModelTag {
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
func (jbobject *ServicesEc2ModelVolume) SetTags(a []*ServicesEc2ModelTag)  {
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

// public com.amazonaws.services.ec2.model.Volume withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelVolume) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Volume withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelVolume) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVolumeType(java.lang.String)
func (jbobject *ServicesEc2ModelVolume) SetVolumeType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVolumeType()
func (jbobject *ServicesEc2ModelVolume) GetVolumeType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolumeType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Volume withVolumeType(java.lang.String)
func (jbobject *ServicesEc2ModelVolume) WithVolumeType2(a string) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeType", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVolumeType(com.amazonaws.services.ec2.model.VolumeType)
func (jbobject *ServicesEc2ModelVolume) SetVolumeType(a ServicesEc2ModelVolumeTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Volume withVolumeType(com.amazonaws.services.ec2.model.VolumeType)
func (jbobject *ServicesEc2ModelVolume) WithVolumeType(a ServicesEc2ModelVolumeTypeInterface) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeType", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeType"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIops(java.lang.Integer)
func (jbobject *ServicesEc2ModelVolume) SetIops(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIops", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getIops()
func (jbobject *ServicesEc2ModelVolume) GetIops() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIops", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.Volume withIops(java.lang.Integer)
func (jbobject *ServicesEc2ModelVolume) WithIops(a int) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIops", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEncrypted(java.lang.Boolean)
func (jbobject *ServicesEc2ModelVolume) SetEncrypted(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEncrypted", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getEncrypted()
func (jbobject *ServicesEc2ModelVolume) GetEncrypted() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEncrypted", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.Volume withEncrypted(java.lang.Boolean)
func (jbobject *ServicesEc2ModelVolume) WithEncrypted(a bool) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEncrypted", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEncrypted()
func (jbobject *ServicesEc2ModelVolume) IsEncrypted() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEncrypted", "java/lang/Boolean")
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

// public void setKmsKeyId(java.lang.String)
func (jbobject *ServicesEc2ModelVolume) SetKmsKeyId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKmsKeyId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getKmsKeyId()
func (jbobject *ServicesEc2ModelVolume) GetKmsKeyId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKmsKeyId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Volume withKmsKeyId(java.lang.String)
func (jbobject *ServicesEc2ModelVolume) WithKmsKeyId(a string) *ServicesEc2ModelVolume {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKmsKeyId", "com/amazonaws/services/ec2/model/Volume", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVolume) ToString() string {
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
func (jbobject *ServicesEc2ModelVolume) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVolume) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.Volume clone()
func (jbobject *ServicesEc2ModelVolume) Clone() *ServicesEc2ModelVolume {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/Volume")
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVolume) Clone2() (*JavaLangObject, error) {
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


