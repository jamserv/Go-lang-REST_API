USE [janezdev]
GO

/****** Object:  Table [dbo].[Users]    Script Date: 07/05/2017 18:34:59 ******/
DROP TABLE [dbo].[Users]
GO

/****** Object:  Table [dbo].[Users]    Script Date: 07/05/2017 18:34:59 ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[Users](
	[uuid] [uniqueidentifier] DEFAULT (NEWID()),
	[name] [varchar](64) NULL,
	[address] [varchar](512) NULL,
	[age] [int] NULL,
	[createDt] [datetime] DEFAULT(GETDATE())
) ON [PRIMARY]

GO