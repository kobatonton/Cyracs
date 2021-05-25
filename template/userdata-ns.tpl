#cloud-config
users:
  - default
  - name: {{.Hostname}}
    gecos: cloud-config user
    groups: users,wheel
    ssh_pwauth: True
chpasswd:
  list: |
    {{.Hostname}}:{{.Hostname}}
  expire: False
write_files:{{range $i, $int := .Interfaces}}
  - path: /etc/sysconfig/network-scripts/ifcfg-if{{$i}}
    permissions: "0644"
    content: |
      HWADDR={{mac $i}}      DEVICE=if{{$i}}
      BOOTPROTO=none
      ONBOOT=yes
      PREFIX={{.Prefix}}
      IPADDR={{.Ip}}{{end -}}
{{if .Packages}}packages:{{end -}}{{range .Packages -}}
  - {{.}}{{end}}
power_state:
  delay: now
  mode: poweroff
  message: Bye Bye
  timeout: 30