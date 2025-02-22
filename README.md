## ⚠️ Disclaimer

This project was built as part of a **hackathon** of Bitcoin++ and, while it has been tested on **Signet**, it is **not production-ready**. Use it at your own risk, and always verify code security before deploying it to mainnet.

This is the **backend** repository for Sats Guardian, built with **Next.js** to provide an intuitive and secure interface for users.

- 🔗 **Frontend Repository:** [Sats Guardian Frontend](https://github.com/gusttav-lang/BtcWatchTower)

- 🏠 **Landing Page:** [Sats Guardian Website](https://hackathon.satsradio.com/)

Check out the backend repository for the core logic and mempool monitoring, and visit our landing page for more details about the project.

## 🛠 Dependencies

To run Sats Guardian, you will need the following dependencies:

- **[Bitcoin Core](https://github.com/bitcoin/bitcoin)** – For monitoring transactions and broadcasting replacements. Compile it with the ZeroMQ flag.

- **[Electrum Server](https://github.com/spesmilo/electrumx)** – For fast UTXO lookups.

- **[BTC Pay Server](https://github.com/btcpayserver/btcpayserver)** – To handle Lightning payments.

- **[Next.js](https://github.com/vercel/next.js)** – The frontend framework for the web interface.

- **[Go](https://github.com/golang/go)** – The backend server, ensuring performance and scalability.

- **[PostgreSQL](https://github.com/postgres/postgres)** – For securely storing user addresses and configurations.

Make sure these dependencies are installed and properly configured before running the project.

## Inspiration

Bitcoin self-custody is powerful, but it comes with risks. If you hold a significant amount of Bitcoin, you need to protect it from threats like:

- **Mnemonic phrase theft** – If someone gains access to your seed, they control your funds.

- **Hot wallet hacks** – Malware or exploits can drain your wallets.

- **The $5 wrench attack** – Physical threats can force you to sign a transaction under duress.

### Do you have time to monitor the mempool 24/7?

Even if you’re aware of these risks, tracking the mempool for potential theft isn’t practical for most individuals. You need an automated, robust solution.

### No need to wait for a soft fork

While **covenant vaults** could offer a long-term solution, waiting for a Bitcoin soft fork isn't an option today. You need protection **now**.

### A simple UX for everyone

Not everyone is a tech expert. You deserve **an easy-to-use** solution that provides top-tier security without complex setups or command-line tools.

### **Introducing Sats Guardian**

Sats Guardian is a **globally distributed network of bots** monitoring the Bitcoin mempool for transactions on your UTXOs.

## What it does?

1.  **Pre-sign recovery transactions** – Before an attack happens, you sign multiple transactions with increasing fees to a **secure recovery address**.

2.  **Automated response** – If a bot detects a robbery transaction, it **instantly broadcasts a pre-signed transaction** with a slightly higher fee, outbidding the attacker.

3.  **Continuous replacement** – The bot keeps replacing the attacker's transaction with a higher-fee transaction until yours gets confirmed.

### Key Features

✅ **No private keys or KYC required** – Maintain full sovereignty over your funds.

✅ **Instant notifications** – Receive alerts via **SMS, email, or Telegram** if an attack attempt is detected.

✅ **Whitelist safe addresses** – Easily spend your UTXOs by **pausing the bot or adding a whitelist** of approved addresses.

✅ **Use it for free (open-source) or via Lightning payments** – Run it yourself or pay a **Lightning invoice** for managed protection.

### **Take control of your Bitcoin security**

Sats Guardian ensures your Bitcoin remains **safe from attackers, without compromising your privacy or sovereignty**.

Protect your funds today!

## How We Built It

Sats Guardian is designed for **speed, reliability, and decentralization**, ensuring that theft detection and response happen as quickly as possible.

### **Global Bitcoin Node Network**

To detect a robbery transaction instantly, we compiled **Bitcoin Core** with the **ZeroMQ flag**, allowing our globally distributed nodes to **publish real-time notifications** to our servers. This ensures that an attack on your UTXOs is detected within seconds.

### **Technology Stack**

🔹 **Frontend** – Built with **Next.js**, providing a fast and intuitive user experience.

🔹 **Backend** – Developed in **Go**, ensuring a high-performance and scalable infrastructure.

🔹 **Bitcoin Data** – We use an **Electrum server** to efficiently search and track user UTXOs.

🔹 **Payments** – Lightning payments are processed through **BTCPay Server**, allowing users to subscribe easily without KYC.

🔹 **Database** – User addresses and configurations are securely stored in **PostgreSQL**.

By combining **distributed Bitcoin nodes, real-time monitoring, and an easy-to-use interface**, Sats Guardian offers **instant protection without compromising decentralization or security**.

## Challenges We Ran Into

Building a **real-time Bitcoin security system** presented several challenges:

🔹 **Mempool latency & network reliability** – Theft transactions propagate quickly, so we needed a **globally distributed** network of Bitcoin nodes to **minimize detection time**.

🔹 **Efficient UTXO tracking** – Searching UTXOs across multiple addresses in a scalable way required integrating **Electrum servers** without sacrificing performance.

🔹 **Handling transaction replacement** – Ensuring the attacker's transaction is **always outbid** without getting stuck in mempool congestion was a major hurdle.

🔹 **Privacy & security** – We designed the system to function **without storing private keys or requiring KYC**, balancing security with user autonomy.

🔹 **Lightning payments** – Implementing **BTCPay Server** for seamless Lightning payments while maintaining a smooth UX was a technical challenge.

Through iteration and testing, we built a **robust and scalable** solution that can safeguard Bitcoin users globally.

## Accomplishments That We're Proud Of

🚀 **Real-time theft detection** – Our globally distributed Bitcoin nodes detect malicious transactions **within seconds**.

🔒 **No private keys, full user sovereignty** – Users remain in complete control of their Bitcoin without trusting a third party.

⚡ **Automated transaction replacement** – Our system **continuously outbids** the attacker's transaction, ensuring your funds are secured.

💡 **Seamless user experience** – We built an intuitive **Next.js frontend** to make Bitcoin security **accessible to everyone**.

🌍 **Decentralized infrastructure** – Multiple independent nodes monitor the Bitcoin network, ensuring a **resilient and censorship-resistant** service.

These accomplishments bring us closer to **making Bitcoin safer for all users**—without waiting for protocol changes or requiring technical expertise.

## What We Learned

Throughout development, we gained valuable insights:

🔹 **Speed matters** – A **few seconds of delay** can determine whether an attack succeeds or fails. Optimizing mempool monitoring was critical.

🔹 **Bitcoin UX needs improvement** – Security solutions must be **simple enough for non-technical users**, or adoption will be limited.

🔹 **Decentralization vs. convenience** – Striking the right balance between **user control and automation** was a key challenge.

We’ve refined our approach based on these lessons, ensuring **Sats Guardian is both powerful and easy to use**.

## What’s Next for Sats Guardian

🔹 **Add auth** – Enter in your account to edit your addresses or stop the bot to create a legit transaction.

🔹 **Mobile app & push notifications** – Get instant alerts via **mobile, SMS, email, or Telegram** when an attack is detected.

🔹 **More configurable security options** – Customize **fee bumping strategies** and recovery addresses for even greater flexibility.

🔹 **Multi-sig & time-locked vaults** – Expanding security options beyond pre-signed transactions for **long-term asset protection**.

🔹 **Tor & privacy enhancements** – Improve privacy by allowing users to interact via **Tor and other anonymity-preserving methods**.

🔹 **Community adoption & audits** – Encourage open-source contributions and **third-party security audits** to ensure trustworthiness.

Sats Guardian is just getting started. Our goal is to **make Bitcoin self-custody safer, without sacrificing decentralization or privacy**. 🚀
