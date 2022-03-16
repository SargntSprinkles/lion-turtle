package playbooks

// func Adamant() Playbook {
// 	return Playbook{
// 		Playbook: "The Adamant",
// 		Principles: []string{
// 			"Restraint",
// 			"Results",
// 		},
// 		startingStats: stats.NewStats(0, 1, -1, 1),
// 		demeanorOptions: []string{
// 			"Above-it-all",
// 			"Perfectionist",
// 			"Chilly",
// 			"Rebellious",
// 			"Flippant",
// 			"Standoffish",
// 		},
// 		historyQuestions: []string{
// 			"What experience of being deceived or manipulated convinced you to steel yourself against being swayed by other people?",
// 			"Who was your first lodestar, and why were they an exception? Why aren't they your lodestar anymore?",
// 			"Who earned your grudging respect by teaching you pragmatism?",
// 			"What heirloom or piece of craftsmanship do you carry to remind you to stay true to yourself?",
// 			"Why are you committed to this group or purpose?",
// 		},
// 		connections: []string{
// 			"__________ takes issue with my methods-perhaps they have a point, but I certianly can't admit that to them!",
// 			"__________ is my lodestar; something about them makes them the one person I let my guard down around.",
// 		},
// 		feature: Feature{
// 			name:        "The Lodestone",
// 			description: `<p>There's only one person you usually let past your emotional walls.</p><p><strong>Name your lodestar</strong> (choose a PC to start):__________</p><p>You can shift your lodestar to someone new when they <strong>guide and comfort</strong> you and you open up to them, or when you <strong>guide and comfort</strong> them and they open up to you. If you do choose to shift your lodestar, clear a condition.</p><p>When you <strong>shut down someone vulnerable to harsh words or icy silence</strong>, shift your balance toward Results and roll with Results. On a hit, they mark a condition and you may clear the same condition. On a 10+, they also cannot shift your balance or <strong>call you out</strong> for the rest of the scene. On a miss, they have exactly the right retort; mark a condition and they shift your balance. You cannot use this on your lodestar.</p><p>When your lodestar <strong>shifts your balance</strong> or <strong>calls you out</strong>, you cannot resist it. Treat an NPC lodestar calling you out as if you rolled a 10+, and a PC lodestar calling you out as if they rolled a 10+.</p><p>When you <strong>consult your lodestar for advice on a problem</strong> (or permission to use your preferred solution), roll with Restraint. On a 10+ take all three; on a 7-9 they choose two:</p><ul><li>You see the wisdom of their advice. They shift your balance; follow their advice and they shift your balance again.</li><li>The conversation bolsters you. Clear a condition or 2-fatigue.</li><li>They feel at ease offering their opinion. They clear a condition or 2-fatigue.</li></ul><p>On a miss, something about their advice infuriates you. Mark a condition or have the GM shift your balance twice.</p>`,
// 		},
// 		moves: []Move{
// 			{
// 				name:        "This Was A Victory",
// 				description: "When you reveal that you have sabotaged a building, device, or vehicle right as it becomes relevant, mark fatigue and roll with <strong>PASSION</strong>. On a hit, your work pays off, creating an opportunity for you and your allies at just the right time. On a 7-9, the opportunity is fleeting—act fast to stay ahead of the consequences. On a miss, your action was ill-judged and something or someone you care about is hurt as collateral damage.",
// 				source:      "The Adamant",
// 			},
// 			{
// 				name:        "Takes One To Know One",
// 				description: "When you verbally needle someone by finding the weaknesses in their armor, roll with Focus. On a hit, ask 1 question. On a 7-9, they ask 1 of you as well:<ul><li>What is your principle?</li><li>What do you need to prove?</li><li>What could shake your certainty?</li><li>Whom do you care about more than you let on?</li></ul>Anyone who lies or stonewalls marks 2-fatigue. On a miss, your attack leaves you exposed; they may ask you any one question from the list, and you must answer honestly.",
// 				source:      "The Adamant",
// 			},
// 			{
// 				name:        "No Time For Feelings",
// 				description: "When you have equal or fewer conditions marked than your highest principle, mark fatigue to push down your feelings for the rest of the scene and ignore condition penalties until the end of the scene. When you <strong>resist an NPC shifting your balance</strong>, mark a condition to roll with conditions marked (max +4). You cannot then choose to clear a condition by immediately proving them wrong.",
// 				source:      "The Adamant",
// 			},
// 			{
// 				name:        "I Don't Hate You",
// 				description: "When you <strong>guide and comfort</strong> someone in an awk- ward, understated, or idiosyncratic fashion, roll with <strong>PASSION</strong> instead of <strong>HARMONY</strong> if you mark Insecure or Insecure is already marked.",
// 				source:      "The Adamant",
// 			},
// 			{
// 				name:        "Driven By Justice",
// 				description: "Take +1 to <strong>PASSION</strong> (max +3).",
// 				source:      "The Adamant",
// 			},
// 		},
// 	}
// }
