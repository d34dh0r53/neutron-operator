package neutronapi

import (
	"github.com/openstack-k8s-operators/lib-common/modules/storage"
	neutronv1beta1 "github.com/openstack-k8s-operators/neutron-operator/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
)

// GetVolumes -
// TODO: merge to GetVolumes when other controllers also switched to current config
//
//	mechanism.
func GetVolumes(name string, extraVol []neutronv1beta1.NeutronExtraVolMounts, svc []storage.PropagationType) []corev1.Volume {
	res := []corev1.Volume{
		{
			Name: "config",
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: name + "-config",
				},
			},
		},
		{
			Name: "httpd-config",
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: name + "-httpd-config",
				},
			},
		},
	}
	for _, exv := range extraVol {
		for _, vol := range exv.Propagate(svc) {
			res = append(res, vol.Volumes...)
		}
	}
	return res

}

// GetVolumeMounts - Neutron API VolumeMounts
func GetVolumeMounts(serviceName string, extraVol []neutronv1beta1.NeutronExtraVolMounts, svc []storage.PropagationType) []corev1.VolumeMount {
	res := []corev1.VolumeMount{
		{
			Name:      "config",
			MountPath: "/var/lib/config-data",
			ReadOnly:  true,
		},
		{
			Name:      "config",
			MountPath: "/var/lib/kolla/config_files/config.json",
			SubPath:   serviceName + "-config.json",
			ReadOnly:  true,
		},
	}
	for _, exv := range extraVol {
		for _, vol := range exv.Propagate(svc) {
			res = append(res, vol.Mounts...)
		}
	}
	return res

} // GetHttpdVolumeMount - Returns the VolumeMounts used by the httpd sidecar
func GetHttpdVolumeMount() []corev1.VolumeMount {
	return []corev1.VolumeMount{
		{
			Name:      "httpd-config",
			MountPath: "/etc/httpd/conf/httpd.conf",
			SubPath:   "httpd.conf",
			ReadOnly:  true,
		},
		{
			Name:      "httpd-config",
			MountPath: "/etc/httpd/conf.d/10-neutron.conf",
			SubPath:   "10-neutron-httpd.conf",
			ReadOnly:  true,
		},
		{
			Name:      "httpd-config",
			MountPath: "/etc/httpd/conf.d/ssl.conf",
			SubPath:   "ssl.conf",
			ReadOnly:  true,
		},
	}
}
