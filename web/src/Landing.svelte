<script lang="ts">
  import Icon from './lib/Icon.svelte'
  import SocialLinks from './lib/SocialLinks.svelte'

  const year = new Date().getFullYear()
</script>

<a class="skip-link" href="#main">Skip to content</a>

<nav class="site-nav" aria-label="Primary">
  <div class="wrap nav-row">
    <a class="logo" href="#main" aria-label="Brainiac home">
      <span class="logo-mark">
        <img src="/assets/img/logo/logo-varied-purple.svg" alt="" width="46" height="37" />
      </span>
      <span class="logo-text">Brainiac</span>
    </a>

    <div class="nav-links">
      <a href="#graphs">Graphs</a>
      <a href="#gamification">Gamification</a>
      <a href="#habits">Habits</a>
      <a href="#create">Create a course</a>
    </div>

    <a class="btn sm nav-cta" href="#">Get started</a>
  </div>
</nav>

<main id="main">
  <section class="hero" aria-labelledby="hero-title">
    <div class="wrap hero-grid">
      <div class="hero-copy">
        <div class="hero-meta">
          <p class="eyebrow">GNU AGPLv3</p>
          <SocialLinks />
        </div>
        <h1 id="hero-title">Brainiac</h1>
        <p class="lead">Graph-first, flexible, gamified.</p>
        <p class="hero-text">
          Brainiac is a free and open source platform that aims to improve the experience of
          students and teachers with courses that use a graph-like structure and gamification
          features.
        </p>
        <div class="btnrow hero-actions">
          <a class="btn" href="#">Get for free</a>
          <a class="btn alt" href="#">Documentation</a>
        </div>
      </div>

      <div class="hero-art">
        <svg
          class="graph"
          viewBox="0 0 560 430"
          role="img"
          aria-label="Course graph with completed Algebra and Statistics, current Functions, upcoming Logic, a checkpoint assessment and locked Geometry and Data modules."
        >
          <path class="e" d="M128 92 C160 92 178 52 210 52" />
          <path class="e lock" d="M298 52 C340 52 350 117 390 117" />
          <path class="e" d="M84 114 C84 170 114 190 114 250" />
          <path class="e" d="M158 272 C185 272 195 352 220 352" />
          <path class="e" d="M308 352 C340 352 355 300 355 254" />
          <path class="e lock" d="M355 210 C355 180 434 175 434 139" />
          <path class="e" d="M158 272 C210 272 250 232 300 232" />
          <path class="e" d="M254 74 C254 140 340 150 355 210" />
          <path class="e lock" d="M410 232 C430 260 444 280 444 310" />
          <path class="e lock" d="M478 117 C510 180 444 220 444 310" />

          <g class="n done">
            <rect x="40" y="70" width="88" height="44" rx="14" />
            <text x="84" y="98">Algebra</text>
          </g>
          <g class="n cur">
            <rect x="210" y="30" width="88" height="44" rx="14" />
            <text x="254" y="58">Functions</text>
          </g>
          <g class="n lock">
            <rect x="390" y="95" width="88" height="44" rx="14" />
            <text x="434" y="123">Geometry</text>
          </g>
          <g class="n att">
            <rect x="300" y="210" width="110" height="44" rx="14" />
            <text x="355" y="238">Assessment</text>
          </g>
          <g class="n done">
            <rect x="70" y="250" width="88" height="44" rx="14" />
            <text x="114" y="278">Statistics</text>
          </g>
          <g class="n">
            <rect x="220" y="330" width="88" height="44" rx="14" />
            <text x="264" y="358">Logic</text>
          </g>
          <g class="n lock">
            <rect x="400" y="310" width="88" height="44" rx="14" />
            <text x="444" y="338">Data</text>
          </g>
        </svg>
      </div>
    </div>
  </section>

  <section id="graphs" class="blk" aria-labelledby="graphs-title">
    <div class="wrap split">
      <div class="copy">
        <p class="eyebrow">Graphs</p>
        <h2 id="graphs-title">See what you know. Choose what is next.</h2>
        <p class="lead">
          Courses in Brainiac have small modules that link to each other in a graph, resembling a
          brain. See clearly what you have already learned and choose topics you want to learn next.
        </p>
        <p>
          Completion may be checked by solving tasks, completing assessments, or manually by
          students. The teacher decides what fits best.
        </p>
        <ul class="methods" aria-label="Ways to complete a module">
          <li><Icon name="check-circle" size={18} /> Task</li>
          <li><Icon name="check-circle" size={18} /> Assessment</li>
          <li><Icon name="check-circle" size={18} /> Manual mark</li>
        </ul>
      </div>

      <div class="demo">
        <div class="panel">
          <div class="panel-head">
            <span class="tag">Course map</span>
            <span class="meta">4 of 7 modules</span>
          </div>
          <h3>Web foundations</h3>
          <div
            class="bar"
            role="progressbar"
            aria-label="Course progress"
            aria-valuenow="57"
            aria-valuemin="0"
            aria-valuemax="100"
          >
            <span style="width: 57%"></span>
          </div>
          <ul class="subjects">
            <li class="subj">
              <span class="dot done" aria-hidden="true"></span>
              <span>HTML basics</span>
              <span class="g">Completed</span>
            </li>
            <li class="subj">
              <span class="dot done" aria-hidden="true"></span>
              <span>CSS layout</span>
              <span class="g">Completed</span>
            </li>
            <li class="subj">
              <span class="dot cur" aria-hidden="true"></span>
              <span>JavaScript</span>
              <span class="g">In progress</span>
            </li>
            <li class="subj">
              <span class="dot lock" aria-hidden="true"></span>
              <span>Accessibility</span>
              <span class="g">Locked</span>
            </li>
            <li class="subj">
              <span class="dot done" aria-hidden="true"></span>
              <span>Checkpoint assessment</span>
              <span class="g">Assessment, 88%</span>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </section>

  <section id="gamification" class="blk surf" aria-labelledby="gamification-title">
    <div class="wrap split flip">
      <div class="copy">
        <p class="eyebrow">Gamification</p>
        <h2 id="gamification-title">Progress that feels like play.</h2>
        <p class="lead">
          By progressing through a course, students level up and earn coins, as well as unlock
          special items and prizes.
        </p>
        <p>
          Teachers can customize what students get, add their own items, and let students exchange
          them for physical gifts.
        </p>
      </div>

      <div class="demo">
        <div class="panel">
          <div class="panel-head">
            <div class="coin">
              <i aria-hidden="true"></i>
              1,240<span class="sr-only"> coins</span>
            </div>
            <span class="tag s">Level 7</span>
          </div>
          <h3>Reward shelf</h3>
          <p class="meta">320 XP to the next level</p>
          <div
            class="bar"
            role="progressbar"
            aria-label="Progress to the next level"
            aria-valuenow="62"
            aria-valuemin="0"
            aria-valuemax="100"
          >
            <span style="width: 62%"></span>
          </div>
          <div class="items">
            <div class="item">
              <div class="art p"><Icon name="bolt" size={64} /></div>
              <b>XP Boost</b>
              <span class="mut">For maintaining daily streak for 5 days</span>
            </div>
            <div class="item">
              <div class="art a"><Icon name="cup-hot" size={64} /></div>
              <b>Coffee token</b>
              <span class="mut">For completing assessment on Java functions</span>
            </div>
            <div class="item">
              <div class="art s"><Icon name="snowflake" size={64} /></div>
              <b>Streak Freeze</b>
              <span class="mut">For reaching level 2. Keeps your streak after a 1 day miss</span>
            </div>
            <div class="item">
              <div class="art p"><Icon name="emoji-funny-square" size={64} /></div>
              <b>Random Sticker</b>
              <span class="mut">Done all modules in CSS. Exchange for a sticker!</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <section id="habits" class="blk" aria-labelledby="habits-title">
    <div class="wrap split">
      <div class="copy">
        <p class="eyebrow">Setting habits</p>
        <h2 id="habits-title">A streak that sticks.</h2>
        <p class="lead">
          Keep a daily learning streak and show it on your profile. It motivates you to learn
          something new every day.
        </p>
        <p>
          Streaks are visible to other members, so a small daily session becomes a habit worth
          keeping.
        </p>
      </div>

      <div class="demo">
        <div class="panel">
          <div class="panel-head">
            <span class="tag">Profile</span>
            <span class="meta">My profile</span>
          </div>
          <div class="profile-top">
            <span class="avatar">
              <img
                src="/assets/img/avatar/crown.svg"
                alt=""
                width="34"
                height="34"
                loading="lazy"
              />
            </span>
            <div>
              <h3>Liza</h3>
              <p class="meta">Learning every day 😉</p>
            </div>
          </div>
          <div class="streak">
            <Icon name="flame" size={30} />
            <b>38</b>
            <span>day streak</span>
          </div>
          <div class="week" role="img" aria-label="Active on 6 of the last 7 days">
            <i class="y"></i>
            <i class="y"></i>
            <i class="y"></i>
            <i class="y"></i>
            <i class="y"></i>
            <i></i>
            <i></i>
          </div>
          <p class="meta week-caption">Last 7 days</p>
          <ul class="facts">
            <li><span>Skills learned</span><b>87</b></li>
            <li><span>Courses finished</span><b>2</b></li>
          </ul>
        </div>
      </div>
    </div>
  </section>

  <section id="create" class="blk surf" aria-labelledby="create-title">
    <div class="wrap split flip">
      <div class="copy">
        <p class="eyebrow">Create a course</p>
        <h2 id="create-title">Write once. Take it anywhere.</h2>
        <p class="lead">
          All elements of a course, including textbooks, tasks and assessments, can be made in the
          integrated web editor.
        </p>
        <p>
          You are not locked into the platform: everything can be exported as Markdown files or in
          other formats, such as PDF.
        </p>
      </div>

      <div class="demo">
        <div
          class="panel"
          role="img"
          aria-label="Course editor mock-up with textbook, tasks and assessment tabs, and Markdown and PDF export options."
        >
          <div class="panel-head">
            <span class="tag">Web editor</span>
            <span class="meta">Autosaved</span>
          </div>
          <div class="editor-tabs" aria-hidden="true">
            <span class="on">Modules</span>
            <span>Tasks</span>
            <span>Assessments</span>
          </div>
          <div class="editor-sheet" aria-hidden="true">
            <span class="line title"></span>
            <span class="line"></span>
            <span class="line short"></span>
            <span class="line"></span>
            <span class="line tiny"></span>
          </div>
          <div class="export-row" aria-hidden="true">
            <span class="export"><Icon name="download" size={18} /> Markdown</span>
            <span class="export"><Icon name="download" size={18} /> PDF</span>
          </div>
        </div>
      </div>
    </div>
  </section>
</main>

<footer class="site-footer">
  <div class="wrap footer-main">
    <div class="footer-info">
      <p>© {year} Brainiac · Free and open source</p>
      <!-- TODO: replace placeholder hrefs with the published legal documents. -->
      <nav class="legal-links" aria-label="Legal">
        <a href="/legal/privacy" title="Coming soon">Privacy policy</a>
        <a href="/legal/terms" title="Coming soon">Terms of service</a>
        <a href="/legal/notice" title="Coming soon">Legal notice</a>
      </nav>
    </div>
    <div class="footer-actions">
      <SocialLinks />
      <a class="footer-top" href="#main">Back to top</a>
    </div>
  </div>
</footer>
