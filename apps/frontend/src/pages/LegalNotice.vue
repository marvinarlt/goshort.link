<script setup lang="ts">
  const {
    VITE_LN_NAME,
    VITE_LN_STREET,
    VITE_LN_POSTAL,
    VITE_LN_CITY,
    VITE_LN_MAIL
  } = import.meta.env;
</script>

<template>
  <h1>{{ $t('navigation.legal-notice.label') }} <span>Angaben gemäß § 5 TMG</span></h1>

  <dl class="styled">
    <dt>Name</dt>
    <dd>{{ VITE_LN_NAME }}</dd>
    
    <dt>Adresse</dt>
    <dd>{{ VITE_LN_STREET }}, {{ VITE_LN_POSTAL }} {{ VITE_LN_CITY }}</dd>

    <dt>E-Mail</dt>
    <dd><a :href="'mailto:' + VITE_LN_MAIL">{{ VITE_LN_MAIL }}</a></dd>
  </dl>

  <h2>Haftungsausschluss</h2>
  
  <h3>Haftung für Links</h3>
  <p>Unser Angebot enthält Links zu externen Webseiten Dritter, auf deren Inhalte wir keinen Einfluss haben. Deshalb können wir für diese fremden Inhalte auch keine Gewähr übernehmen. Für die Inhalte der verlinkten Seiten ist stets der jeweilige Anbieter oder Betreiber der Seiten verantwortlich. Die verlinkten Seiten wurden zum Zeitpunkt der Verlinkung auf mögliche Rechtsverstöße überprüft. Rechtswidrige Inhalte waren zum Zeitpunkt der Verlinkung nicht erkennbar. Eine permanente inhaltliche Kontrolle der verlinkten Seiten ist jedoch ohne konkrete Anhaltspunkte einer Rechtsverletzung nicht zumutbar. Bei Bekanntwerden von Rechtsverletzungen werden wir derartige Links umgehend entfernen.</p>

  <h3>Urheberrecht</h3>
  <p> Die durch die Seitenbetreiber erstellten Inhalte und Werke auf diesen Seiten unterliegen dem deutschen Urheberrecht. Die Vervielfältigung, Bearbeitung, Verbreitung und jede Art der Verwertung außerhalb der Grenzen des Urheberrechtes bedürfen der schriftlichen Zustimmung des jeweiligen Autors bzw. Erstellers. Downloads und Kopien dieser Seite sind nur für den privaten, nicht kommerziellen Gebrauch gestattet. Soweit die Inhalte auf dieser Seite nicht vom Betreiber erstellt wurden, werden die Urheberrechte Dritter beachtet. Insbesondere werden Inhalte Dritter als solche gekennzeichnet. Sollten Sie trotzdem auf eine Urheberrechtsverletzung aufmerksam werden, bitten wir um einen entsprechenden Hinweis. Bei Bekanntwerden von Rechtsverletzungen werden wir derartige Inhalte umgehend entfernen.</p>

  <h3>Datenschutz</h3>
  <p>Der Nutzung von im Rahmen der Impressumspflicht veröffentlichten Kontaktdaten durch Dritte zur Übersendung von nicht ausdrücklich angeforderter Werbung und Informationsmaterialien wird hiermit ausdrücklich widersprochen. Die Betreiber der Seiten behalten sich ausdrücklich rechtliche Schritte im Falle der unverlangten Zusendung von Werbeinformationen, etwa durch Spam-Mails, vor.</p>

  <p>Website Impressum erstellt durch <a href="https://www.impressum-generator.de" target="_blank">impressum-generator.de</a> von der <a href="https://www.kanzlei-hasselbach.de/" rel="nofollow" target="_blank">Kanzlei Hasselbach</a></p>
</template>