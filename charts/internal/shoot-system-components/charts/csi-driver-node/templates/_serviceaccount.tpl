{{- define "csi-driver-node.serviceaccount" -}}
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: csi-driver-node-{{ .role }}
  namespace: {{ .Release.Namespace }}
{{- if .Values.global.useProjectedTokenMount }}
automountServiceAccountToken: false
{{- end }}
{{- end -}}
