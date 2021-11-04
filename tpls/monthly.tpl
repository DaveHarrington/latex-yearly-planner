\begin{tiny}
\begin{tabularx}{\textwidth}{@{}
  {{- if $.Cfg.Pages.WeeklyEnabled -}}
    l!{\vrule width \myLenLineThicknessThick}
  {{- else -}} | {{- end -}}
  *{7}{@{}X@{}|}}
\noalign{\hrule height \myLenLineThicknessDefault}
{{.Body.WeekHeaderFull $.Cfg.Pages.WeeklyEnabled}} \\ \noalign{\hrule height \myLenLineThicknessDefault}

{{ .Body.MatrixTexed .Cfg.Pages.WeeklyEnabled true true nil }}

\end{tabularx}
\end{tiny}
\medskip

\parbox{\myLenTwoCol}{
  \myUnderline{Notes}
  \vbox to \dimexpr\textheight-\pagetotal-\myLenLineHeightButLine\relax {%
    \leaders\hbox to \linewidth{\textcolor{\myColorGray}{\rule{0pt}{\myLenLineHeightButLine}\hrulefill}}\vfil
  }%
}%
\hspace{\myLenTwoColSep}%
\parbox{\myLenTwoCol}{
    \myUnderline{Notes}
    \vbox to \dimexpr\textheight-\pagetotal-\myLenLineHeightButLine\relax {%
        \leaders\hbox to \linewidth{\textcolor{\myColorGray}{\rule{0pt}{\myLenLineHeightButLine}\hrulefill}}\vfil
    }%
}

\pagebreak
