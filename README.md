# latex-yearly-planner

PDF planner designed for e-ink devices.

See [discussions](https://github.com/kudrykv/latex-yearly-planner/discussions) for available planners and their variations.

# Preview examples
<img src="https://github.com/kudrykv/latex-yearly-planner/blob/main/examples/pictures/sn_a5x.planner/01_annual.png" width="419"><img src="https://github.com/kudrykv/latex-yearly-planner/blob/main/examples/pictures/sn_a5x.planner/02_quarter.png" width="419"><img src="https://github.com/kudrykv/latex-yearly-planner/blob/main/examples/pictures/sn_a5x.planner/03_month.png" width="419"><img src="https://github.com/kudrykv/latex-yearly-planner/blob/main/examples/pictures/sn_a5x.planner/04_week.png" width="419"><img src="https://github.com/kudrykv/latex-yearly-planner/blob/main/examples/pictures/sn_a5x.planner/05_day.png" width="419"><img src="https://github.com/kudrykv/latex-yearly-planner/blob/main/examples/pictures/sn_a5x.planner/06_day_notes.png" width="419"><img src="https://github.com/kudrykv/latex-yearly-planner/blob/main/examples/pictures/sn_a5x.planner/07_day_reflect.png" width="419"><img src="https://github.com/kudrykv/latex-yearly-planner/blob/main/examples/pictures/sn_a5x.planner/08_todos_index.png" width="419"><img src="https://github.com/kudrykv/latex-yearly-planner/blob/main/examples/pictures/sn_a5x.planner/09_todos_page.png" width="419"><img src="https://github.com/kudrykv/latex-yearly-planner/blob/main/examples/pictures/sn_a5x.planner/10_notes_index.png" width="419"><img src="https://github.com/kudrykv/latex-yearly-planner/blob/main/examples/pictures/sn_a5x.planner/11_notes_page.png" width="419">

# Dave notes

`docker run -it --name planner ubuntu:latest bash`

Then get back:
`docker exec -it planner bash`

--
git checkout https://github.com/DaveHarrington/latex-yearly-planner
cd latex-yearly-planner
git checkout dave-planner
./release.sh

CONT=`docker ls | grep planner | cut -d' ' -f1` # untested :P
docker cp $CONT:/proj/latex-yearly-planner/2021.rm2.planner.pdf . && open *.pdf

