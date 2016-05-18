package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceCapacityInterface interface {
	JavaLangObjectInterface

	// public void setInstanceType(java.lang.String)
	SetInstanceType(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.InstanceCapacity withInstanceType(java.lang.String)
	WithInstanceType(a string) *ServicesEc2ModelInstanceCapacity

	// public void setAvailableCapacity(java.lang.Integer)
	SetAvailableCapacity(a int) 

	// public java.lang.Integer getAvailableCapacity()
	GetAvailableCapacity() int

	// public com.amazonaws.services.ec2.model.InstanceCapacity withAvailableCapacity(java.lang.Integer)
	WithAvailableCapacity(a int) *ServicesEc2ModelInstanceCapacity

	// public void setTotalCapacity(java.lang.Integer)
	SetTotalCapacity(a int) 

	// public java.lang.Integer getTotalCapacity()
	GetTotalCapacity() int

	// public com.amazonaws.services.ec2.model.InstanceCapacity withTotalCapacity(java.lang.Integer)
	WithTotalCapacity(a int) *ServicesEc2ModelInstanceCapacity

	// public com.amazonaws.services.ec2.model.InstanceCapacity clone()
	Clone() *ServicesEc2ModelInstanceCapacity
}

type ServicesEc2ModelInstanceCapacity struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceCapacity()
func NewServicesEc2ModelInstanceCapacity() (*ServicesEc2ModelInstanceCapacity) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceCapacity")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceCapacity{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceCapacity) SetInstanceType(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceType()
func (jbobject *ServicesEc2ModelInstanceCapacity) GetInstanceType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceCapacity withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceCapacity) WithInstanceType(a string) *ServicesEc2ModelInstanceCapacity {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/InstanceCapacity", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceCapacity{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailableCapacity(java.lang.Integer)
func (jbobject *ServicesEc2ModelInstanceCapacity) SetAvailableCapacity(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailableCapacity", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getAvailableCapacity()
func (jbobject *ServicesEc2ModelInstanceCapacity) GetAvailableCapacity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailableCapacity", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.InstanceCapacity withAvailableCapacity(java.lang.Integer)
func (jbobject *ServicesEc2ModelInstanceCapacity) WithAvailableCapacity(a int) *ServicesEc2ModelInstanceCapacity {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailableCapacity", "com/amazonaws/services/ec2/model/InstanceCapacity", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelInstanceCapacity{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTotalCapacity(java.lang.Integer)
func (jbobject *ServicesEc2ModelInstanceCapacity) SetTotalCapacity(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTotalCapacity", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getTotalCapacity()
func (jbobject *ServicesEc2ModelInstanceCapacity) GetTotalCapacity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTotalCapacity", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.InstanceCapacity withTotalCapacity(java.lang.Integer)
func (jbobject *ServicesEc2ModelInstanceCapacity) WithTotalCapacity(a int) *ServicesEc2ModelInstanceCapacity {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTotalCapacity", "com/amazonaws/services/ec2/model/InstanceCapacity", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelInstanceCapacity{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceCapacity) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceCapacity) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceCapacity) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceCapacity clone()
func (jbobject *ServicesEc2ModelInstanceCapacity) Clone() *ServicesEc2ModelInstanceCapacity {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceCapacity")
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
	unique_x := &ServicesEc2ModelInstanceCapacity{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceCapacity) Clone2() (*JavaLangObject, error) {
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


