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
write_files:
  - path: /etc/netplan/50-cloud-init.yaml
    permissions: "0644"
    content: |
      network:
        version: 2
        ethernets:
          def:
            match:
              macaddress: {{defmac}}            wakeonlan: true
            dhcp4: true{{range $i, $int := .Interfaces}}
          if{{$i}}:
            match:
              macaddress: {{mac $i}}            wakeonlan: true
            addresses:
              - {{.Ip}}/{{.Prefix -}}
            {{if .Gateway}}gateway4: {{.Gateway}}{{end}}{{end -}}
{{if .Packages}}packages:{{end -}}{{range .Packages -}}
  - {{.}}{{end}}
power_state:
  delay: now
  mode: poweroff
  message: Bye Bye
  timeout: 30