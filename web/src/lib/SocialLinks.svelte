<script lang="ts">
  import SocialIcon from './SocialIcon.svelte'
</script>

<div class="social-links" role="group" aria-label="Project links">
  <a
    class="social-link"
    href="https://github.com/ReallyWeirdCat/brainiac"
    target="_blank"
    rel="noopener noreferrer"
    aria-label="Brainiac on GitHub"
  >
    <SocialIcon name="github" size={22} />
  </a>
  <a
    class="social-link"
    href="https://t.me/project_brainiac"
    target="_blank"
    rel="noopener noreferrer"
    aria-label="Brainiac on Telegram"
  >
    <SocialIcon name="telegram" size={22} />
  </a>
  <!-- TODO: replace with the project Discord invite link. -->
  <a
    class="social-link"
    href="/discord"
    title="Discord invite coming soon"
    aria-label="Brainiac on Discord"
  >
    <SocialIcon name="discord" size={22} />
  </a>
  <!-- TODO: add the project contact address. -->
  <a class="social-link" href="mailto:" aria-label="Email Brainiac" title="Email address coming soon">
    <SocialIcon name="email" size={22} />
  </a>
</div>
