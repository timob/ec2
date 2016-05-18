package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSpotFleetRequestHistoryResultInterface interface {
	JavaLangObjectInterface

	// public void setSpotFleetRequestId(java.lang.String)
	SetSpotFleetRequestId(a string) 

	// public java.lang.String getSpotFleetRequestId()
	GetSpotFleetRequestId() string

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult withSpotFleetRequestId(java.lang.String)
	WithSpotFleetRequestId(a string) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult

	// public void setStartTime(java.util.Date)
	SetStartTime(a time.Time) 

	// public java.util.Date getStartTime()
	GetStartTime() time.Time

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult withStartTime(java.util.Date)
	WithStartTime(a time.Time) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult

	// public void setLastEvaluatedTime(java.util.Date)
	SetLastEvaluatedTime(a time.Time) 

	// public java.util.Date getLastEvaluatedTime()
	GetLastEvaluatedTime() time.Time

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult withLastEvaluatedTime(java.util.Date)
	WithLastEvaluatedTime(a time.Time) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult

	// public java.util.List<com.amazonaws.services.ec2.model.HistoryRecord> getHistoryRecords()
	GetHistoryRecords() []*ServicesEc2ModelHistoryRecord

	// public void setHistoryRecords(java.util.Collection<com.amazonaws.services.ec2.model.HistoryRecord>)
	SetHistoryRecords(a []*ServicesEc2ModelHistoryRecord) 

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult withHistoryRecords(com.amazonaws.services.ec2.model.HistoryRecord...)
	WithHistoryRecords(a ...*ServicesEc2ModelHistoryRecord) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult withHistoryRecords(java.util.Collection<com.amazonaws.services.ec2.model.HistoryRecord>)
	WithHistoryRecords2(a []*ServicesEc2ModelHistoryRecord) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult clone()
	Clone() *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult
}

type ServicesEc2ModelDescribeSpotFleetRequestHistoryResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult()
func NewServicesEc2ModelDescribeSpotFleetRequestHistoryResult() (*ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSpotFleetRequestHistoryResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSpotFleetRequestHistoryResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSpotFleetRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) SetSpotFleetRequestId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotFleetRequestId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSpotFleetRequestId()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) GetSpotFleetRequestId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotFleetRequestId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult withSpotFleetRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) WithSpotFleetRequestId(a string) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestId", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestHistoryResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStartTime(java.util.Date)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) SetStartTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStartTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getStartTime()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) GetStartTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStartTime", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult withStartTime(java.util.Date)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) WithStartTime(a time.Time) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStartTime", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestHistoryResult", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLastEvaluatedTime(java.util.Date)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) SetLastEvaluatedTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLastEvaluatedTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getLastEvaluatedTime()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) GetLastEvaluatedTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLastEvaluatedTime", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult withLastEvaluatedTime(java.util.Date)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) WithLastEvaluatedTime(a time.Time) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLastEvaluatedTime", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestHistoryResult", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.HistoryRecord> getHistoryRecords()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) GetHistoryRecords() []*ServicesEc2ModelHistoryRecord {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHistoryRecords", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelHistoryRecord)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setHistoryRecords(java.util.Collection<com.amazonaws.services.ec2.model.HistoryRecord>)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) SetHistoryRecords(a []*ServicesEc2ModelHistoryRecord)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHistoryRecords", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult withHistoryRecords(com.amazonaws.services.ec2.model.HistoryRecord...)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) WithHistoryRecords(a ...*ServicesEc2ModelHistoryRecord) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/HistoryRecord")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHistoryRecords", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestHistoryResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/HistoryRecord"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult withHistoryRecords(java.util.Collection<com.amazonaws.services.ec2.model.HistoryRecord>)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) WithHistoryRecords2(a []*ServicesEc2ModelHistoryRecord) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHistoryRecords", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestHistoryResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) SetNextToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNextToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNextToken()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) GetNextToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNextToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) WithNextToken(a string) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestHistoryResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult clone()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) Clone() *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestHistoryResult")
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult) Clone2() (*JavaLangObject, error) {
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


