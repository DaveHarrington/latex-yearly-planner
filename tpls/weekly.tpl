{{- $week := .Body -}}
{{- $day1 := index $week 0 -}}
{{- $day2 := index $week 1 -}}
{{- $day3 := index $week 2 -}}
{{- $day4 := index $week 3 -}}
{{- $day5 := index $week 4 -}}
{{- $day6 := index $week 5 -}}
{{- $day7 := index $week 6 -}}

\parbox{\myLenTwoCol}{
  \myUnderline{
    \hyperlink{ {{- $day1.RefText -}} }{ {{$day1.Weekday -}} } |
    \hyperlink{ {{- $day2.RefText -}} }{ {{$day2.Weekday -}} } |
    \hyperlink{ {{- $day3.RefText -}} }{ {{$day3.Weekday -}} }
  }
  \vbox to \dimexpr\textheight-\pagetotal-\myLenLineHeightButLine\relax {%
    \leaders\hbox to \linewidth{\textcolor{\myColorGray}{\rule{0pt}{\myLenLineHeightButLine}\hrulefill}}\vfil
  }%
}%
\hspace{\myLenTwoColSep}%
\parbox{\myLenTwoCol}{
  \myUnderline{
    \hyperlink{ {{- $day4.RefText -}} }{ {{$day4.Weekday -}} } |
    \hyperlink{ {{- $day5.RefText -}} }{ {{$day5.Weekday -}} }
  }
  \vbox to \dimexpr\textheight-\pagetotal-\myLenLineHeightButLine\relax {%
    \leaders\hbox to \linewidth{\textcolor{\myColorGray}{\rule{0pt}{\myLenLineHeightButLine}\hrulefill}}\vfil
  }%
}

\pagebreak
