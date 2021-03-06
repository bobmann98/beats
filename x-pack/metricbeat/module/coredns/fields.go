// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package coredns

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "coredns", asset.ModuleFieldsPri, AssetCoredns); err != nil {
		panic(err)
	}
}

// AssetCoredns returns asset data.
// This is the base64 encoded gzipped contents of module/coredns.
func AssetCoredns() string {
	return "eJzElk9T2zwQxu/5FM/kBO+A32mnwyEz7QUO5VDKAPeiSJtExZaMVoYJn74j/0kc28SEqoNOnpX07E/rR2uf4oHWM0jrSBmeAF77lGaY1pHpBFDE0unca2tm+DYB0KzHD6uKlCaAo5QE0wxLMQEWmlLFs3LpKYzIqJ0iDL/Ow2Jni7yODGQJ477edw9pjRfaMNgLr9lryfAr4fFMjuBIKCyczcpEF1e3tUIbpY0TRHgTHQLaAxXG+QCOo1R4UvAWfkUNCZjck5bU2r5brmZ0Wdu8uTBaJtIWxu/MN+SpNcvOxB74ctxZL1KYIpuTg11UKXgwuzKcOHosiP2/YHgsyK3RFx5KrwonglxiOJkX8oF88t8gjZ3/JtkFrYK/3s17U1GgocBKs7dLJzJUMAxtYISxTNIaNV7O9nm4yN5a2IV1mfCzDcj7zsGb/SfgIgsuaAJ/dZCoHhlg3Zq2zj6Ox/qFkvna00fa5la/UN8yh7AfYJERqFeYuglGeGK+6VeI3tgXbFzXXW0sFrqTpvpjsxJPVDb3i5+Yaw+mcbTAEBcO50FtFy8nhxdrCMKoUnmQqzfRUDzQ+tk6dRjId5sqLstRtfAgFZhCpD79nupwbg1T4qRVcevTqk6TpVOfMucgWn8mUoVK4W1xKqxBhIXIdLqOx3C3IgilHDHX2g2Gd8Jwbp3H0Sd8xeU1ji6v8USOw7fty/EJPpfhs5342fHx+Fv90H6Lmxrj8I7bx4/Zcoep9vfcPlHcpjvM9HrXDb+z5OLas9JsroWepxT+cEb7SLjPcUGCYtBGwaS6DP/vvba5s97Gpdnez07fwNG0UPkU1mHqZT7dXsj+AzqOkkKuKFlpH9dI50EWQbbyzqZ4ZT7kabHUZgQq08yR/V1hVcL7wP4EAAD//4GDEiY="
}
