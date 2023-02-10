{{- define "app.labels" -}}
app: {{ .Release.Name }}
{{- end -}}

{{- define "cloudrun.labels" -}}
{{- if (eq .Values.type "cloudrun") -}}
networking.knative.dev/visibility: cluster-local
{{- end -}}
{{- end -}}

{{- define "labels" -}}
labels:
{{- include "app.labels" . | nindent 2 -}}
{{- end -}}

{{- define "cloudrun.annotations" -}}
{{- if and ((.Values.hpa).enabled) (eq .Values.type "cloudrun") -}}
autoscaling.knative.dev/minScale: {{ .Values.hpa.minReplicas | quote }}
autoscaling.knative.dev/maxScale: {{ .Values.hpa.maxReplicas | quote }}
{{- else if (.Values.replicas) -}}
autoscaling.knative.dev/minScale: {{ .Values.replicas | quote }}
{{- end -}}
{{- end -}}

{{- define "annotations" -}}
annotations:
{{- include "cloudrun.annotations" . | nindent 2 -}}
{{- end -}}

{{- define "metadata" -}}
metadata:
{{- include "labels" . | nindent 2 -}}
{{- include "annotations" . | nindent 2 -}}
{{- end -}}

{{- define "volume.logging" -}}
- name: logging
  configMap:
    name: {{ .Release.Name }}-logging
{{- end -}}

{{- define "cloudsql" -}}
{{- if eq ((.Values.cloudsql).enabled) true -}}
- name: cloud-sql-proxy
  image: gcr.io/cloudsql-docker/gce-proxy:1.28.0
  command:
    - "/cloud_sql_proxy"
    # Replace DB_PORT with the port the proxy should listen on
    # Defaults: MySQL: 3306, Postgres: 5432, SQLServer: 1433
    - "-instances={{ required "Cloud SQL instance connection must be provided." .Values.cloudsql.instances }}"
    - "-enable_iam_login"
  securityContext:
    runAsNonRoot: true
  resources:
    requests:
      memory: "500Mi"
      cpu: "300m"
    limits:
      memory: "500Mi"
      cpu: "300m"
{{- end -}}
{{- end -}}

{{- define "typeCheck" -}}
{{- if not (or (eq ((.Values).type) "kubernetes") (eq ((.Values).type) "cloudrun")) -}}
{{- fail "Type must be one of: kubernetes, cloudrun" -}} 
{{- end -}}
{{- end -}}

{{- define "service.ports" -}}
{{- if not (has (float64 80) ((.Values.service).ports)) -}}
- name: "80"
  port: 80
  targetPort: {{ $.Values.application.containerPort }}
{{- end -}}
{{- range $port := ((.Values.service).ports) }}
- name: {{ $port | quote }}
  port: {{ $port }}
  targetPort: {{ $.Values.application.containerPort }}
{{- end }}
{{- end -}}

{{- define "serviceAccount" -}}
{{- if ne .Values.serviceAccountName "" -}}
serviceAccountName: {{ .Values.serviceAccountName }}
{{- end -}}
{{- end -}}