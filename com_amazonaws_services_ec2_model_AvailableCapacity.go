package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAvailableCapacityInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceCapacity> getAvailableInstanceCapacity()
	GetAvailableInstanceCapacity() []*ServicesEc2ModelInstanceCapacity

	// public void setAvailableInstanceCapacity(java.util.Collection<com.amazonaws.services.ec2.model.InstanceCapacity>)
	SetAvailableInstanceCapacity(a []*ServicesEc2ModelInstanceCapacity) 

	// public com.amazonaws.services.ec2.model.AvailableCapacity withAvailableInstanceCapacity(com.amazonaws.services.ec2.model.InstanceCapacity...)
	WithAvailableInstanceCapacity(a ...*ServicesEc2ModelInstanceCapacity) *ServicesEc2ModelAvailableCapacity

	// public com.amazonaws.services.ec2.model.AvailableCapacity withAvailableInstanceCapacity(java.util.Collection<com.amazonaws.services.ec2.model.InstanceCapacity>)
	WithAvailableInstanceCapacity2(a []*ServicesEc2ModelInstanceCapacity) *ServicesEc2ModelAvailableCapacity

	// public void setAvailableVCpus(java.lang.Integer)
	SetAvailableVCpus(a int) 

	// public java.lang.Integer getAvailableVCpus()
	GetAvailableVCpus() int

	// public com.amazonaws.services.ec2.model.AvailableCapacity withAvailableVCpus(java.lang.Integer)
	WithAvailableVCpus(a int) *ServicesEc2ModelAvailableCapacity

	// public com.amazonaws.services.ec2.model.AvailableCapacity clone()
	Clone() *ServicesEc2ModelAvailableCapacity
}

type ServicesEc2ModelAvailableCapacity struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.AvailableCapacity()
func NewServicesEc2ModelAvailableCapacity() (*ServicesEc2ModelAvailableCapacity) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AvailableCapacity")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAvailableCapacity{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceCapacity> getAvailableInstanceCapacity()
func (jbobject *ServicesEc2ModelAvailableCapacity) GetAvailableInstanceCapacity() []*ServicesEc2ModelInstanceCapacity {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailableInstanceCapacity", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInstanceCapacity)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAvailableInstanceCapacity(java.util.Collection<com.amazonaws.services.ec2.model.InstanceCapacity>)
func (jbobject *ServicesEc2ModelAvailableCapacity) SetAvailableInstanceCapacity(a []*ServicesEc2ModelInstanceCapacity)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailableInstanceCapacity", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.AvailableCapacity withAvailableInstanceCapacity(com.amazonaws.services.ec2.model.InstanceCapacity...)
func (jbobject *ServicesEc2ModelAvailableCapacity) WithAvailableInstanceCapacity(a ...*ServicesEc2ModelInstanceCapacity) *ServicesEc2ModelAvailableCapacity {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceCapacity")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailableInstanceCapacity", "com/amazonaws/services/ec2/model/AvailableCapacity", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceCapacity"))
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
	unique_x := &ServicesEc2ModelAvailableCapacity{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AvailableCapacity withAvailableInstanceCapacity(java.util.Collection<com.amazonaws.services.ec2.model.InstanceCapacity>)
func (jbobject *ServicesEc2ModelAvailableCapacity) WithAvailableInstanceCapacity2(a []*ServicesEc2ModelInstanceCapacity) *ServicesEc2ModelAvailableCapacity {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailableInstanceCapacity", "com/amazonaws/services/ec2/model/AvailableCapacity", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelAvailableCapacity{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailableVCpus(java.lang.Integer)
func (jbobject *ServicesEc2ModelAvailableCapacity) SetAvailableVCpus(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailableVCpus", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getAvailableVCpus()
func (jbobject *ServicesEc2ModelAvailableCapacity) GetAvailableVCpus() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailableVCpus", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.AvailableCapacity withAvailableVCpus(java.lang.Integer)
func (jbobject *ServicesEc2ModelAvailableCapacity) WithAvailableVCpus(a int) *ServicesEc2ModelAvailableCapacity {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailableVCpus", "com/amazonaws/services/ec2/model/AvailableCapacity", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelAvailableCapacity{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAvailableCapacity) ToString() string {
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
func (jbobject *ServicesEc2ModelAvailableCapacity) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAvailableCapacity) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AvailableCapacity clone()
func (jbobject *ServicesEc2ModelAvailableCapacity) Clone() *ServicesEc2ModelAvailableCapacity {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AvailableCapacity")
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
	unique_x := &ServicesEc2ModelAvailableCapacity{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelAvailableCapacity) Clone2() (*JavaLangObject, error) {
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


