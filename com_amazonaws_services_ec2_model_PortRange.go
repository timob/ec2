package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPortRangeInterface interface {
	JavaLangObjectInterface

	// public void setFrom(java.lang.Integer)
	SetFrom(a int) 

	// public java.lang.Integer getFrom()
	GetFrom() int

	// public com.amazonaws.services.ec2.model.PortRange withFrom(java.lang.Integer)
	WithFrom(a int) *ServicesEc2ModelPortRange

	// public void setTo(java.lang.Integer)
	SetTo(a int) 

	// public java.lang.Integer getTo()
	GetTo() int

	// public com.amazonaws.services.ec2.model.PortRange withTo(java.lang.Integer)
	WithTo(a int) *ServicesEc2ModelPortRange

	// public com.amazonaws.services.ec2.model.PortRange clone()
	Clone() *ServicesEc2ModelPortRange
}

type ServicesEc2ModelPortRange struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.PortRange()
func NewServicesEc2ModelPortRange() (*ServicesEc2ModelPortRange) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PortRange")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPortRange{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setFrom(java.lang.Integer)
func (jbobject *ServicesEc2ModelPortRange) SetFrom(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFrom", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getFrom()
func (jbobject *ServicesEc2ModelPortRange) GetFrom() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFrom", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.PortRange withFrom(java.lang.Integer)
func (jbobject *ServicesEc2ModelPortRange) WithFrom(a int) *ServicesEc2ModelPortRange {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFrom", "com/amazonaws/services/ec2/model/PortRange", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelPortRange{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTo(java.lang.Integer)
func (jbobject *ServicesEc2ModelPortRange) SetTo(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTo", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getTo()
func (jbobject *ServicesEc2ModelPortRange) GetTo() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTo", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.PortRange withTo(java.lang.Integer)
func (jbobject *ServicesEc2ModelPortRange) WithTo(a int) *ServicesEc2ModelPortRange {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTo", "com/amazonaws/services/ec2/model/PortRange", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelPortRange{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPortRange) ToString() string {
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
func (jbobject *ServicesEc2ModelPortRange) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelPortRange) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.PortRange clone()
func (jbobject *ServicesEc2ModelPortRange) Clone() *ServicesEc2ModelPortRange {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/PortRange")
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
	unique_x := &ServicesEc2ModelPortRange{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelPortRange) Clone2() (*JavaLangObject, error) {
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


