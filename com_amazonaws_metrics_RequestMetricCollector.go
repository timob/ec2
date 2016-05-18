package ec2

import "github.com/timob/javabind"

type MetricsRequestMetricCollectorInterface interface {
	JavaLangObjectInterface

	// public abstract void collectMetrics(com.amazonaws.Request<?>, com.amazonaws.Response<?>)
	CollectMetrics(a RequestInterface, b ResponseInterface) 

	// public boolean isEnabled()
	IsEnabled() bool
}

type MetricsRequestMetricCollector struct {
	JavaLangObject
}

// public com.amazonaws.metrics.RequestMetricCollector()
func NewMetricsRequestMetricCollector() (*MetricsRequestMetricCollector) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/metrics/RequestMetricCollector")
	if err != nil {
		panic(err)
	}
	x := &MetricsRequestMetricCollector{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public abstract void collectMetrics(com.amazonaws.Request<?>, com.amazonaws.Response<?>)
func (jbobject *MetricsRequestMetricCollector) CollectMetrics(a RequestInterface, b ResponseInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "collectMetrics", javabind.Void, conv_a.Value().Cast("com/amazonaws/Request"), conv_b.Value().Cast("com/amazonaws/Response"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public boolean isEnabled()
func (jbobject *MetricsRequestMetricCollector) IsEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *MetricsRequestMetricCollector) NONE() *MetricsRequestMetricCollector {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/metrics/RequestMetricCollector", "NONE", "com/amazonaws/metrics/RequestMetricCollector")
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
	unique_x := &MetricsRequestMetricCollector{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *MetricsRequestMetricCollector) SetFieldNONE(val MetricsRequestMetricCollectorInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/metrics/RequestMetricCollector", "NONE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


