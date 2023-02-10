{{- define "app.name" -}}
- name: {{ .Release.Name }}
{{- end -}}

{{- define "app.image" -}}
{{- if eq .Values.application.image.tag "latest" -}}
{{- fail "Cannot use latest tag!" -}}
{{- else -}}
image: {{ printf "%s:%s" ( required "Application image must be specified." .Values.application.image.repository ) ( required "Application tag must be specified." .Values.application.image.tag ) }}
{{- end }}
{{- end -}}

{{- define "app.ports" -}}
ports:
  - containerPort: {{ required "Container port must be provided." .Values.application.containerPort }}
{{- end -}}


{{- define "app.env.custom" -}}
{{- range $var := .Values.application.env }}
- name: {{ $var.name }}
  value: {{ $var.value | quote }}
{{- end -}}
{{- end -}}

{{- define "app.env" -}}
env:
{{- include "app.env.custom" . | nindent 2 -}}
{{- end -}}

{{- define "app.probe.startup" -}}
{{- if (and (.Values.application.startupProbe) (ne .Values.type "cloudrun" ) ) -}}
startupProbe:
  httpGet:
    path: {{ required "Startup probe health check path must be provided." ((.Values.application.startupProbe).httpPath) }}
    port: {{ .Values.application.containerPort }}
  failureThreshold: {{ ((.Values.application.startupProbe).failureThreshold) | default 10 }}
  periodSeconds: {{ ((.Values.application.startupProbe).period) | default 10 }}
{{- end -}}
{{- end -}}

{{- define "app.probe.liveness" -}}
{{- if (.Values.application.livenessProbe) -}}
livenessProbe:
  httpGet:
    path: {{ required "Liveness probe health check path must be provided." .Values.application.livenessProbe.httpPath }}
    port: {{ .Values.application.containerPort }}
  periodSeconds: {{ .Values.application.livenessProbe.periodSeconds | default 5 }}
{{- end -}}
{{- end -}}

{{- define "app.probe.readiness" -}}
{{- if (.Values.application.readinessProbe) -}}
readinessProbe:
  httpGet:
    path: {{ required "Readiness probe health check path must be provided." .Values.application.readinessProbe.httpPath }}
    port: {{ .Values.application.containerPort }}
  periodSeconds: {{ .Values.application.readinessProbe.periodSeconds | default 5 }}
{{- end -}}
{{- end -}}

{{- define "app.resources" -}}
{{- if ((.Values.application).resources) -}}
resources:
  {{- if .Values.application.resources.requests }}
  requests:
    {{- if .Values.application.resources.requests.memory }}
    memory: {{ .Values.application.resources.requests.memory }}
    {{- end -}}
    {{- if .Values.application.resources.requests.cpu }}
    cpu: {{ .Values.application.resources.requests.cpu }}
    {{- end -}}
  {{- end -}}
  {{- if .Values.application.resources.limits }}
  limits:
    {{- if .Values.application.resources.limits.memory }}
    memory: {{ .Values.application.resources.limits.memory }}
    {{- end -}}
    {{- if .Values.application.resources.limits.cpu }}
    cpu: {{ .Values.application.resources.limits.cpu }}
    {{- end -}}
  {{- end -}}
{{- end -}}
{{- end -}}

{{- define "app.volumeMount.logging" -}}
- name: logging
  mountPath: "/config"
  readOnly: true
{{- end -}}

{{- define "app" -}}
{{- include "cloudsql" . | nindent 0 -}}
{{- include "app.name" . | nindent 0 -}}
{{- include "app.image" . | nindent 2 -}}
{{- include "app.ports" . | nindent 2 -}}
{{- include "app.env" . | nindent 2 -}}
{{- include "app.probe.startup" . | nindent 2 -}}
{{- include "app.probe.liveness" . | nindent 2 -}}
{{- include "app.probe.readiness" . | nindent 2 -}}
{{- include "app.resources" . | nindent 2 }}
  imagePullPolicy: Always
  volumeMounts:
{{- include "app.volumeMount.logging" . | nindent 4 -}}
{{- end -}}