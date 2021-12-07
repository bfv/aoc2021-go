define variable timings as int64 extent 9 no-undo.

function init returns logical () forward.
function solve returns int (days as int) forward.
function shift returns logical() forward.

etime(true).
init().
message solve(256) etime(false) "ms" view-as alert-box.

function solve returns int (days as int):
  define variable answer as int64 no-undo.
  define variable i as integer no-undo.
  do i = 1 to days:
    shift().
  end.
  do i = 1 to 9:
    answer = answer + timings[i].
  end.
  return answer.
end.

function shift returns logical():
  define variable j as integer no-undo.
  define variable newFish as int64 no-undo.
  newFish = timings[1].
  do j = 1 to 8:
    timings[j] = timings[j + 1].
  end.
  timings[7] += newFish.
  timings[9] = newFish.   
end.

function init returns logical ():
  define variable inputString as character no-undo.
  define variable numEntries as integer no-undo.
  define variable i as integer no-undo.
  define variable ix as integer no-undo.
  input from value(search("input.txt")).  
  import unformatted inputString.
  numEntries = num-entries(inputString).
  do i = 1 to numEntries:
    ix = int(entry(i, inputString)) + 1.
    timings[ix] = timings[ix] + 1.
  end.
end.
