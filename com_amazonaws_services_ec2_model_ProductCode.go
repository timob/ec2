package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelProductCodeInterface interface {
	JavaLangObjectInterface

	// public void setProductCodeId(java.lang.String)
	SetProductCodeId(a string) 

	// public java.lang.String getProductCodeId()
	GetProductCodeId() string

	// public com.amazonaws.services.ec2.model.ProductCode withProductCodeId(java.lang.String)
	WithProductCodeId(a string) *ServicesEc2ModelProductCode

	// public void setProductCodeType(java.lang.String)
	SetProductCodeType2(a string) 

	// public java.lang.String getProductCodeType()
	GetProductCodeType() string

	// public com.amazonaws.services.ec2.model.ProductCode withProductCodeType(java.lang.String)
	WithProductCodeType2(a string) *ServicesEc2ModelProductCode

	// public void setProductCodeType(com.amazonaws.services.ec2.model.ProductCodeValues)
	SetProductCodeType(a ServicesEc2ModelProductCodeValuesInterface) 

	// public com.amazonaws.services.ec2.model.ProductCode withProductCodeType(com.amazonaws.services.ec2.model.ProductCodeValues)
	WithProductCodeType(a ServicesEc2ModelProductCodeValuesInterface) *ServicesEc2ModelProductCode

	// public com.amazonaws.services.ec2.model.ProductCode clone()
	Clone() *ServicesEc2ModelProductCode
}

type ServicesEc2ModelProductCode struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ProductCode()
func NewServicesEc2ModelProductCode() (*ServicesEc2ModelProductCode) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ProductCode")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelProductCode{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setProductCodeId(java.lang.String)
func (jbobject *ServicesEc2ModelProductCode) SetProductCodeId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductCodeId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getProductCodeId()
func (jbobject *ServicesEc2ModelProductCode) GetProductCodeId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProductCodeId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ProductCode withProductCodeId(java.lang.String)
func (jbobject *ServicesEc2ModelProductCode) WithProductCodeId(a string) *ServicesEc2ModelProductCode {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodeId", "com/amazonaws/services/ec2/model/ProductCode", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelProductCode{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProductCodeType(java.lang.String)
func (jbobject *ServicesEc2ModelProductCode) SetProductCodeType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductCodeType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getProductCodeType()
func (jbobject *ServicesEc2ModelProductCode) GetProductCodeType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProductCodeType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ProductCode withProductCodeType(java.lang.String)
func (jbobject *ServicesEc2ModelProductCode) WithProductCodeType2(a string) *ServicesEc2ModelProductCode {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodeType", "com/amazonaws/services/ec2/model/ProductCode", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelProductCode{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProductCodeType(com.amazonaws.services.ec2.model.ProductCodeValues)
func (jbobject *ServicesEc2ModelProductCode) SetProductCodeType(a ServicesEc2ModelProductCodeValuesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductCodeType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ProductCodeValues"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ProductCode withProductCodeType(com.amazonaws.services.ec2.model.ProductCodeValues)
func (jbobject *ServicesEc2ModelProductCode) WithProductCodeType(a ServicesEc2ModelProductCodeValuesInterface) *ServicesEc2ModelProductCode {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodeType", "com/amazonaws/services/ec2/model/ProductCode", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ProductCodeValues"))
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
	unique_x := &ServicesEc2ModelProductCode{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelProductCode) ToString() string {
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
func (jbobject *ServicesEc2ModelProductCode) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelProductCode) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ProductCode clone()
func (jbobject *ServicesEc2ModelProductCode) Clone() *ServicesEc2ModelProductCode {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ProductCode")
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
	unique_x := &ServicesEc2ModelProductCode{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelProductCode) Clone2() (*JavaLangObject, error) {
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


