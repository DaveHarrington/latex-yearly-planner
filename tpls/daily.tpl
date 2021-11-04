{{- $today := .Body.Day -}}

\myUnderline{To Do\textcolor{white}{g}\hfill{}}
\myTodoneLineGray{Review weekly plan}
\myTodoneLineGray{Block time for: exercise, stretch, meditate}
\myTodoneLineGray{Calendar: today \& tomorrow}
\myTodoneLineGray{Todoist}
\Repeat{\myNumDailyTodos}{\myTodoLineGray}
\vskip\dimexpr7mm-4pt
\myUnderline{Notes $\vert$ \hyperlink{notes {{- $today.RefText -}} }{More}\hfill{}\hfill{}\hyperlink{Notes Index}{All notes}}
\Repeat{\myNumDailyNotes}{\myLineGrayVskipTop}

\pagebreak
