package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRegionInterface interface {
	JavaLangObjectInterface

	// public void setRegionName(java.lang.String)
	SetRegionName(a string) 

	// public java.lang.String getRegionName()
	GetRegionName() string

	// public com.amazonaws.services.ec2.model.Region withRegionName(java.lang.String)
	WithRegionName(a string) *ServicesEc2ModelRegion

	// public void setEndpoint(java.lang.String)
	SetEndpoint(a string) 

	// public java.lang.String getEndpoint()
	GetEndpoint() string

	// public com.amazonaws.services.ec2.model.Region withEndpoint(java.lang.String)
	WithEndpoint(a string) *ServicesEc2ModelRegion

	// public com.amazonaws.services.ec2.model.Region clone()
	Clone() *ServicesEc2ModelRegion
}

type ServicesEc2ModelRegion struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.Region()
func NewServicesEc2ModelRegion() (*ServicesEc2ModelRegion) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/Region")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRegion{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setRegionName(java.lang.String)
func (jbobject *ServicesEc2ModelRegion) SetRegionName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRegionName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRegionName()
func (jbobject *ServicesEc2ModelRegion) GetRegionName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRegionName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Region withRegionName(java.lang.String)
func (jbobject *ServicesEc2ModelRegion) WithRegionName(a string) *ServicesEc2ModelRegion {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRegionName", "com/amazonaws/services/ec2/model/Region", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRegion{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEndpoint(java.lang.String)
func (jbobject *ServicesEc2ModelRegion) SetEndpoint(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEndpoint", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getEndpoint()
func (jbobject *ServicesEc2ModelRegion) GetEndpoint() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEndpoint", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Region withEndpoint(java.lang.String)
func (jbobject *ServicesEc2ModelRegion) WithEndpoint(a string) *ServicesEc2ModelRegion {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEndpoint", "com/amazonaws/services/ec2/model/Region", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRegion{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRegion) ToString() string {
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
func (jbobject *ServicesEc2ModelRegion) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelRegion) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.Region clone()
func (jbobject *ServicesEc2ModelRegion) Clone() *ServicesEc2ModelRegion {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/Region")
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
	unique_x := &ServicesEc2ModelRegion{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelRegion) Clone2() (*JavaLangObject, error) {
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


