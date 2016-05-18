package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelScheduledInstancesEbsInterface interface {
	JavaLangObjectInterface

	// public void setSnapshotId(java.lang.String)
	SetSnapshotId(a string) 

	// public java.lang.String getSnapshotId()
	GetSnapshotId() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs withSnapshotId(java.lang.String)
	WithSnapshotId(a string) *ServicesEc2ModelScheduledInstancesEbs

	// public void setVolumeSize(java.lang.Integer)
	SetVolumeSize(a int) 

	// public java.lang.Integer getVolumeSize()
	GetVolumeSize() int

	// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs withVolumeSize(java.lang.Integer)
	WithVolumeSize(a int) *ServicesEc2ModelScheduledInstancesEbs

	// public void setDeleteOnTermination(java.lang.Boolean)
	SetDeleteOnTermination(a bool) 

	// public java.lang.Boolean getDeleteOnTermination()
	GetDeleteOnTermination() bool

	// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs withDeleteOnTermination(java.lang.Boolean)
	WithDeleteOnTermination(a bool) *ServicesEc2ModelScheduledInstancesEbs

	// public java.lang.Boolean isDeleteOnTermination()
	IsDeleteOnTermination() bool

	// public void setVolumeType(java.lang.String)
	SetVolumeType(a string) 

	// public java.lang.String getVolumeType()
	GetVolumeType() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs withVolumeType(java.lang.String)
	WithVolumeType(a string) *ServicesEc2ModelScheduledInstancesEbs

	// public void setIops(java.lang.Integer)
	SetIops(a int) 

	// public java.lang.Integer getIops()
	GetIops() int

	// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs withIops(java.lang.Integer)
	WithIops(a int) *ServicesEc2ModelScheduledInstancesEbs

	// public void setEncrypted(java.lang.Boolean)
	SetEncrypted(a bool) 

	// public java.lang.Boolean getEncrypted()
	GetEncrypted() bool

	// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs withEncrypted(java.lang.Boolean)
	WithEncrypted(a bool) *ServicesEc2ModelScheduledInstancesEbs

	// public java.lang.Boolean isEncrypted()
	IsEncrypted() bool

	// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs clone()
	Clone() *ServicesEc2ModelScheduledInstancesEbs
}

type ServicesEc2ModelScheduledInstancesEbs struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs()
func NewServicesEc2ModelScheduledInstancesEbs() (*ServicesEc2ModelScheduledInstancesEbs) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ScheduledInstancesEbs")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelScheduledInstancesEbs{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) SetSnapshotId(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) GetSnapshotId() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs withSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) WithSnapshotId(a string) *ServicesEc2ModelScheduledInstancesEbs {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotId", "com/amazonaws/services/ec2/model/ScheduledInstancesEbs", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesEbs{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVolumeSize(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) SetVolumeSize(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeSize", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getVolumeSize()
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) GetVolumeSize() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolumeSize", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs withVolumeSize(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) WithVolumeSize(a int) *ServicesEc2ModelScheduledInstancesEbs {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeSize", "com/amazonaws/services/ec2/model/ScheduledInstancesEbs", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesEbs{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeleteOnTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) SetDeleteOnTermination(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDeleteOnTermination", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getDeleteOnTermination()
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) GetDeleteOnTermination() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeleteOnTermination", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs withDeleteOnTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) WithDeleteOnTermination(a bool) *ServicesEc2ModelScheduledInstancesEbs {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeleteOnTermination", "com/amazonaws/services/ec2/model/ScheduledInstancesEbs", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesEbs{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isDeleteOnTermination()
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) IsDeleteOnTermination() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDeleteOnTermination", "java/lang/Boolean")
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

// public void setVolumeType(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) SetVolumeType(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) GetVolumeType() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs withVolumeType(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) WithVolumeType(a string) *ServicesEc2ModelScheduledInstancesEbs {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeType", "com/amazonaws/services/ec2/model/ScheduledInstancesEbs", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesEbs{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIops(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) SetIops(a int)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) GetIops() int {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs withIops(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) WithIops(a int) *ServicesEc2ModelScheduledInstancesEbs {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIops", "com/amazonaws/services/ec2/model/ScheduledInstancesEbs", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesEbs{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEncrypted(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) SetEncrypted(a bool)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) GetEncrypted() bool {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs withEncrypted(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) WithEncrypted(a bool) *ServicesEc2ModelScheduledInstancesEbs {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEncrypted", "com/amazonaws/services/ec2/model/ScheduledInstancesEbs", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesEbs{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEncrypted()
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) IsEncrypted() bool {
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

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) ToString() string {
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
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs clone()
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) Clone() *ServicesEc2ModelScheduledInstancesEbs {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ScheduledInstancesEbs")
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
	unique_x := &ServicesEc2ModelScheduledInstancesEbs{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelScheduledInstancesEbs) Clone2() (*JavaLangObject, error) {
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


