package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPricingDetailInterface interface {
	JavaLangObjectInterface

	// public void setPrice(java.lang.Double)
	SetPrice(a float64) 

	// public java.lang.Double getPrice()
	GetPrice() float64

	// public com.amazonaws.services.ec2.model.PricingDetail withPrice(java.lang.Double)
	WithPrice(a float64) *ServicesEc2ModelPricingDetail

	// public void setCount(java.lang.Integer)
	SetCount(a int) 

	// public java.lang.Integer getCount()
	GetCount() int

	// public com.amazonaws.services.ec2.model.PricingDetail withCount(java.lang.Integer)
	WithCount(a int) *ServicesEc2ModelPricingDetail

	// public com.amazonaws.services.ec2.model.PricingDetail clone()
	Clone() *ServicesEc2ModelPricingDetail
}

type ServicesEc2ModelPricingDetail struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.PricingDetail()
func NewServicesEc2ModelPricingDetail() (*ServicesEc2ModelPricingDetail) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PricingDetail")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPricingDetail{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setPrice(java.lang.Double)
func (jbobject *ServicesEc2ModelPricingDetail) SetPrice(a float64)  {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrice", javabind.Void, conv_a.Value().Cast("java/lang/Double"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Double getPrice()
func (jbobject *ServicesEc2ModelPricingDetail) GetPrice() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrice", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.PricingDetail withPrice(java.lang.Double)
func (jbobject *ServicesEc2ModelPricingDetail) WithPrice(a float64) *ServicesEc2ModelPricingDetail {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrice", "com/amazonaws/services/ec2/model/PricingDetail", conv_a.Value().Cast("java/lang/Double"))
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
	unique_x := &ServicesEc2ModelPricingDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelPricingDetail) SetCount(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCount", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getCount()
func (jbobject *ServicesEc2ModelPricingDetail) GetCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCount", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.PricingDetail withCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelPricingDetail) WithCount(a int) *ServicesEc2ModelPricingDetail {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCount", "com/amazonaws/services/ec2/model/PricingDetail", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelPricingDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPricingDetail) ToString() string {
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
func (jbobject *ServicesEc2ModelPricingDetail) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelPricingDetail) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.PricingDetail clone()
func (jbobject *ServicesEc2ModelPricingDetail) Clone() *ServicesEc2ModelPricingDetail {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/PricingDetail")
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
	unique_x := &ServicesEc2ModelPricingDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelPricingDetail) Clone2() (*JavaLangObject, error) {
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


