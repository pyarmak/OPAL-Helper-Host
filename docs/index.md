---
layout: default
---

# Notice
> You **MUST** have the host application installed on your computer in order to use the browser extension.

OPAL Helper was designed to aid medical students at the University of Manitoba
by augmenting the [Online Portal for Advanced Learning](https://opal.med.umanitoba.ca).

# Motivation
One faithful autumn evening in 2020, while avoiding watching lecture videos,
I got fed up with OPAL's inexplicable usage of the deprecated VLC npapi plugin.
So, as any (un)reasonable medical student would, I went ahead and wrote a
[Google chrome](https://www.google.com/chrome/browser/desktop/index.html)
extension and accompanying host application in order to launch session videos in
an external video player (_with speed controls!_).

In my initial release over a year ago, the only thing OPAL Helper was able to do was to launch our session videos
in the external player of our choosing and storing a history of recently viewed session videos
for easy access through the extension's popup.

# Features
- Allows streaming session videos using an external player.
    - Comes with a built in [mpv player](https://mpv.io/).
- Allows downloading of session videos for offline viewing.
- Provides a "Recently Viewed" popup (accessed by clicking the extension icon)
    - Clicking on any of the rows will take you to the session's main page.
    - Clicking on the play icon on the right will launch the session video directly.

By the way, you can also go to the OPAL homepage by clicking on the title of the popup
(OPAL Helper).

# Future Development

This was initially conceived as a weekend project which I re-wrote this year for macOS compatibility so if you
find any bugs please submit them to the [issue tracker]({{ site.github.issues_url }}).

#### Happy studying!
