### There should be some data shown here!


#### 💻 Hackatime Stats
{{with hackatimeStats}}
- Status: {{.Data.Status}}
- Total Time: {{.Data.HumanReadableTotal}}
- Daily Average: {{.Data.HumanReadableDailyAvg}}
{{range .Data.Languages}}
  - {{.Name}}: {{.Text}} ({{.Percent}}%)
{{end}}
{{end}}



#### 🤲🤲🤲🤲🤲🤲🤲

{{with hackatimeStats}}
{{ wakatimeLanguages "💾 Languages:" .Data.Languages 5 .Data.HumanReadableTotal }}
{{end}}

