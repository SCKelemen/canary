module github.com/SCKelemen/canary

go 1.13

require (
	exclaim v0.0.0
	token v0.0.0
	yell v0.0.0
)

replace exclaim v0.0.0 => ./exclaim

replace yell v0.0.0 => ./yell

replace token v0.0.0 => ./token
