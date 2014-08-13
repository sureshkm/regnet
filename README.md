regnet 
======
Status : Work In Progress (prototype phase)


Complex regular expressions from simple manageable regex units

Example:

```
DAY = `(?:Mon(?:day)?|Tue(?:sday)?|Wed(?:nesday)?|Thu(?:rsday)?|Fri(?:day)?|Sat(?:urday)?|Sun(?:day)?)`
```

```
MONTH = `(?:Jan(?:uary)?|Feb(?:ruary)?|Mar(?:ch)?|Apr(?:il)?|May|Jun(?:e)?|Jul(?:y)?|Aug(?:ust)?|Sep(?:tember)?|Oct(?:ober)?|Nov(?:ember)?|Dec(?:ember)?)`
```

```
YEAR = `(\d\d){1,2}`
```

%{DAY} %{MONTH} %{YEAR}

can match a text **"The date on which this note is written is Tue Aug 2014"** to give **"Tue Aug 2014"**
