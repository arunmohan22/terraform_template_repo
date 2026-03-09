{# raw tag tells jinja to start escaping double curly brackets #}{%- raw -%}
{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "{% endraw %}{{cookiecutter.serviceSlug}}.name" -}}{% raw %}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "{% endraw %}{{cookiecutter.serviceSlug}}.fullname" -}}{% raw %}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "{% endraw %}{{cookiecutter.serviceSlug}}{% raw %}.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "{% endraw %}{{cookiecutter.serviceSlug}}{% raw %}.labels" -}}
helm.sh/chart: {{ include "{% endraw %}{{cookiecutter.serviceSlug}}{% raw %}.chart" . }}
{{ include "{% endraw %}{{cookiecutter.serviceSlug}}{% raw %}.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "{% endraw %}{{cookiecutter.serviceSlug}}{% raw %}.selectorLabels" -}}
app.kubernetes.io/name: {{ include "{% endraw %}{{cookiecutter.serviceSlug}}{% raw %}.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "{% endraw %}{{cookiecutter.serviceSlug}}{% raw %}.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "{% endraw %}{{cookiecutter.serviceSlug}}{% raw %}.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/* Allows us to specify environment variables entirely within the values file(s) */}}
{{- define "{% endraw %}{{cookiecutter.serviceSlug}}{% raw %}.valueOrValueFrom" -}}
        {{- if kindIs "map" . }}
valueFrom:
  configMapKeyRef:
    name: {{ required "configMapName (the name of the config map to fetch env var from) must be specified." .configMapName | quote }}
    key: {{ required "configMapKey (the key inside the config map to fetch env var from) must be specified." .configMapKey | quote }}
        {{- else }}
value: {{ . | quote -}}
    {{- end }}
{{- end }}
{%- endraw -%}{# endraw tag tells jinja to stop escaping double curly brackets #}
